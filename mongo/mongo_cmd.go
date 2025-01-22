package mongo

import (
	"context"

	"github.com/adiatma85/own-go-sdk/instrument"
	"github.com/adiatma85/own-go-sdk/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Command interface {
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
	InsertOne(ctx context.Context, collectionName string, name string, item interface{}) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, collectionName string, name string, filter interface{}) (*mongo.SingleResult, error)
	Find(ctx context.Context, collectionName, name string, filter interface{}, findOptions *options.FindOptions) (*mongo.Cursor, error)
	CountDocument(ctx context.Context, collectionName, name string, filter interface{}) (int64, error)
	Update(ctx context.Context, collectionName string, name string, filter, updatedItem interface{}) (*mongo.UpdateResult, error)
	Upsert(ctx context.Context, collectionName, name string, filter, insertItem interface{}) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, collectionName string, name string, filter interface{}) (*mongo.DeleteResult, error)
}

type Option struct {
	DisableLimit bool `form:"disableLimit"`
	IsActive     bool
	IsInactive   bool
}

// Note: Kind of dumb, this wrapper will be specific to mongodb for now
// idk about something like sql package that can connect to various provider like `mysql` and `postgres`
type command struct {
	client        *mongo.Client
	database      *mongo.Database
	connName      string
	connType      string
	log           log.Interface
	instrument    instrument.Interface
	useInstrument bool
}

func initCommand(client *mongo.Client, cfg Config, instr instrument.Interface, log log.Interface, isLeader bool) Command {
	conf := cfg.Follower
	connType := connTypeFollower
	if isLeader {
		conf = cfg.Leader
		connType = connTypeLeader
	}

	// Initialize database
	database := client.Database(conf.DB)

	return &command{
		client:        client,
		database:      database,
		connName:      cfg.Name,
		connType:      connType,
		log:           log,
		instrument:    instr,
		useInstrument: cfg.UseInstrument,
	}
}

func (c *command) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}

func (c *command) Ping(ctx context.Context) error {
	return c.client.Ping(ctx, nil)
}

func (c *command) InsertOne(ctx context.Context, collectionName, name string, item interface{}) (*mongo.InsertOneResult, error) {
	return c.database.Collection(collectionName).InsertOne(ctx, item)
}

// Find One
func (c *command) FindOne(ctx context.Context, collectionName, name string, filter interface{}) (*mongo.SingleResult, error) {
	singleResult := c.database.Collection(collectionName).FindOne(ctx, filter)

	return singleResult, singleResult.Err()
}

// Find Multiple
func (c *command) Find(ctx context.Context, collectionName, name string, filter interface{}, findOptions *options.FindOptions) (*mongo.Cursor, error) {
	return c.database.Collection(collectionName).Find(ctx, filter, findOptions)

}

func (c *command) CountDocument(ctx context.Context, collectionName, name string, filter interface{}) (int64, error) {
	return c.database.Collection(collectionName).CountDocuments(ctx, filter)
}

// Update
func (c *command) Update(ctx context.Context, collectionName, name string, filter, updatedItem interface{}) (*mongo.UpdateResult, error) {
	return c.database.Collection(collectionName).UpdateMany(ctx, filter, updatedItem)
}

// Upsert
func (c *command) Upsert(ctx context.Context, collectionName, name string, filter, insertItem interface{}) (*mongo.UpdateResult, error) {
	upsertOptions := options.Update().SetUpsert(true)
	return c.database.Collection(collectionName).UpdateMany(ctx, filter, insertItem, upsertOptions)

}

// Delete
func (c *command) Delete(ctx context.Context, collectionName, name string, filter interface{}) (*mongo.DeleteResult, error) {
	return c.database.Collection(collectionName).DeleteMany(ctx, filter)
}
