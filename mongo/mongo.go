package mongo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/adiatma85/own-go-sdk/instrument"
	"github.com/adiatma85/own-go-sdk/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	connTypeLeader   = "leader"
	connTypeFollower = "follower"
)

type Config struct {
	UseInstrument bool
	LogQuery      bool
	Driver        string
	Name          string
	Enable        bool
	Leader        ConnConfig
	Follower      ConnConfig
}

type ConnConfig struct {
	Host     string
	Port     int
	DB       string
	User     string
	Password string
	TLS      TLSOption
	Options  ConnOptions
}

type TLSOption struct {
	Enable             bool
	CertificateKeyFile string
	KeyFile            string
	CAFile             string
	Insecure           bool
}

type ConnOptions struct {
	AuthSource  string
	MaxLifeTime time.Duration
	MaxIdle     int
	MaxOpen     int
}

type Interface interface {
	Leader() Command
	Follower() Command
	Stop()
}

type documentDB struct {
	endOnce    *sync.Once
	follower   Command
	leader     Command
	cfg        Config
	log        log.Interface
	instrument instrument.Interface
}

// TODO: Coverage
func Init(cfg Config, log log.Interface, instr instrument.Interface) Interface {
	docDB := &documentDB{
		endOnce:    &sync.Once{},
		log:        log,
		cfg:        cfg,
		instrument: instr,
	}

	// If enabled, then do the init
	if cfg.Enable {
		docDB.initDB()
	}

	return docDB
}

func (d *documentDB) Leader() Command {
	return d.leader
}

func (d *documentDB) Follower() Command {
	return d.follower
}

func (d *documentDB) Stop() {
	d.endOnce.Do(func() {
		ctx := context.Background()
		if d.leader != nil {
			if err := d.leader.Close(ctx); err != nil {
				d.log.Error(ctx, err)
			}
		}

		if d.follower != nil {
			if err := d.follower.Close(ctx); err != nil {
				d.log.Error(ctx, err)
			}
		}
	})
}

func (d *documentDB) initDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := d.connect(true, ctx)
	if err != nil {
		d.log.Fatal(ctx, fmt.Sprintf("[FATAL] cannot connect to db %s leader: %s on port %d, with error: %s", d.cfg.Leader.DB, d.cfg.Leader.Host, d.cfg.Leader.Port, err))
	}

	d.log.Info(ctx, fmt.Sprintf("MONGODB: [LEADER] driver=%s db=%s @%s:%v ssl=%v", d.cfg.Driver, d.cfg.Leader.DB, d.cfg.Leader.Host, d.cfg.Leader.Port, d.cfg.Leader.TLS.Enable))
	d.leader = initCommand(db, d.cfg, d.instrument, d.log, true)

	if d.isFollowerEnabled() {
		db, err := d.connect(false, ctx)
		if err != nil {
			d.log.Fatal(ctx, fmt.Sprintf("[FATAL] cannot connect to db %s follower: %s on port %d, with error: %s", d.cfg.Follower.DB, d.cfg.Follower.Host, d.cfg.Follower.Port, err))

			d.log.Info(ctx, fmt.Sprintf("MONGODB: [FOLLOWER] driver=%s db=%s @%s:%v ssl=%v", d.cfg.Driver, d.cfg.Follower.DB, d.cfg.Follower.Host, d.cfg.Follower.Port, d.cfg.Follower.TLS.Enable))
			d.follower = initCommand(db, d.cfg, d.instrument, d.log, false)
		}
	} else {
		d.follower = d.leader
	}
}

// Dia return mongo client dan error
func (d *documentDB) connect(toLeader bool, ctx context.Context) (*mongo.Client, error) {
	conf := d.cfg.Leader
	if !toLeader {
		conf = d.cfg.Follower
	}

	uri, err := d.getURI(conf)
	if err != nil {
		return nil, err
	}

	// Construct the options
	clientOptions, err := d.buildOptions(uri, conf)
	if err != nil {
		return nil, err
	}

	// Connect to the Mongo
	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return db, nil
}

func (d *documentDB) isFollowerEnabled() bool {
	isHostNotEmpty := d.cfg.Follower.Host != ""
	isHostDifferent := (d.cfg.Follower.Host != d.cfg.Leader.Host && d.cfg.Follower.Port == d.cfg.Leader.Port)
	isPortDifferent := (d.cfg.Follower.Host == d.cfg.Leader.Host && d.cfg.Follower.Port != d.cfg.Leader.Port)
	return isHostNotEmpty && (isHostDifferent || isPortDifferent)
}

func (d *documentDB) buildOptions(uri string, conf ConnConfig) (*options.ClientOptions, error) {
	clientOptions := options.Client().ApplyURI(uri).
		SetMaxConnIdleTime(conf.Options.MaxLifeTime).
		SetMaxPoolSize(uint64(conf.Options.MaxOpen)).
		SetMinPoolSize(uint64(conf.Options.MaxIdle)).
		SetMonitor(d.setupCommandMonitor())

	// If we are using TLS
	if conf.TLS.Enable {

		// Load CA file
		caCert, err := os.ReadFile(conf.TLS.CAFile)
		if err != nil {
			return nil, fmt.Errorf("could not read CA file: %v", err)
		}

		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
			return nil, fmt.Errorf("failed to append CA certificates")
		}

		// Create tls.Config with the loaded CA and client certificates
		tlsConfig := &tls.Config{
			RootCAs:            caCertPool,
			InsecureSkipVerify: conf.TLS.Insecure, // Set to true if you want to skip verification (not recommended)
		}

		// Of ypu are using CertificateKeyFile
		if conf.TLS.CertificateKeyFile != "" && conf.TLS.KeyFile != "" {
			cert, err := tls.LoadX509KeyPair(conf.TLS.CertificateKeyFile, conf.TLS.KeyFile)
			if err != nil {
				return nil, fmt.Errorf("could not load client key pair: %v", err)
			}

			tlsConfig.Certificates = []tls.Certificate{cert}
		}

		clientOptions = clientOptions.SetTLSConfig(tlsConfig)
	}

	return clientOptions, nil
}

func (d *documentDB) getURI(conf ConnConfig) (string, error) {
	switch d.cfg.Driver {
	case "mongodb":
		credentials := ""
		if conf.User != "" && conf.Password != "" {
			credentials = fmt.Sprintf("%s:%s@", conf.User, conf.Password)
		}

		// Default to authSource=admin
		authSource := "?authSource=admin"
		if conf.Options.AuthSource != "" {
			authSource = fmt.Sprintf("?authSource=%s", conf.Options.AuthSource)
		}

		return fmt.Sprintf("mongodb://%s%s:%d/%s%s", credentials, conf.Host, conf.Port, conf.DB, authSource), nil
	default:
		return "", fmt.Errorf(`DB Driver [%s] is not supported`, d.cfg.Driver)
	}
}
