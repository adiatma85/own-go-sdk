package sql

import (
	"testing"

	"github.com/adiatma85/own-go-sdk/instrument"
	"github.com/adiatma85/own-go-sdk/log"
	mock_log "github.com/adiatma85/own-go-sdk/tests/mock/log"
	"github.com/golang/mock/gomock"
)

func initTestDatabase() Interface {
	return Init(Config{
		UseInstrument: true,
		LogQuery:      true,
		Driver:        "mysql",
		Leader: ConnConfig{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "delos_db",
			User:     "root",
			Password: "password",
		},
		Follower: ConnConfig{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "delos_db",
			User:     "root",
			Password: "password",
		},
	}, log.Init(log.Config{Level: "debug"}),
		instrument.Init(instrument.Config{
			Metrics: instrument.MetricsConfig{
				Enabled: true,
			},
		}))
}

func TestInit(t *testing.T) {
	ctrl := gomock.NewController(t)
	logger := mock_log.NewMockInterface(ctrl)
	logger.EXPECT().Error(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Fatal(gomock.Any(), gomock.Any()).AnyTimes()
	logger.EXPECT().Info(gomock.Any(), gomock.Any()).AnyTimes()

	conConfig := ConnConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		DB:       "delos_db",
		User:     "root",
		Password: "password",
	}
	logInit := log.Init(log.Config{Level: "debug"})
	instr := instrument.Init(instrument.Config{
		Metrics: instrument.MetricsConfig{
			Enabled: true,
		},
	})

	type args struct {
		cfg   Config
		log   log.Interface
		instr instrument.Interface
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "no driver name",
			args: args{
				cfg: Config{
					UseInstrument: true,
					LogQuery:      true,
					Driver:        "",
					Leader:        conConfig,
					Follower:      conConfig,
				},
				log:   logger,
				instr: instr,
			},
		},
		{
			name: "ok",
			args: args{
				cfg: Config{
					UseInstrument: true,
					LogQuery:      true,
					Driver:        "mysql",
					Leader:        conConfig,
					Follower:      conConfig,
				},
				log:   logInit,
				instr: instr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.cfg, tt.args.log, tt.args.instr)
		})
	}
}

func TestLeader(t *testing.T) {
	mockDb := initTestDatabase()

	tests := []struct {
		name string
	}{
		{
			name: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDb.Leader()
		})
	}
}

func TestFollower(t *testing.T) {
	mockDb := initTestDatabase()

	tests := []struct {
		name string
	}{
		{
			name: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDb.Follower()
		})
	}
}

func TestStop(t *testing.T) {
	mockDb := initTestDatabase()

	tests := []struct {
		name string
	}{
		{
			name: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDb.Stop()
		})
	}
}
