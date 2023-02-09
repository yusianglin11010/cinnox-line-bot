package database

import (
	"context"
	"fmt"

	"github.com/yusianglin11010/cinnox-line-bot/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	*mongo.Client
}

var mongoClient Mongo

// MongoDB connection
func Initialize(config *config.MongoConfig) {
	mongoClient.initialize(config)
}

func Close() {
	mongoClient.close()
}

func (m *Mongo) initialize(cfg *config.MongoConfig) {
	ctx := context.Background()
	uri := config.MongoURI(cfg)
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(uri))
	if err != nil {
		panic(fmt.Sprint("fail to connect to mongo: ", err))
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	m.Client = client
}

func GetMongo() *Mongo {
	return &mongoClient
}

func (m *Mongo) close() {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (m *Mongo) InitLineMessage(collName string) error {
	db := m.Client.Database("message")
	ctx := context.Background()

	collections, err := db.ListCollectionNames(ctx, bson.D{{}})
	if err != nil {
		return fmt.Errorf("failed to list collection names: %s", err.Error())
	}
	for _, name := range collections {
		if name == collName {
			return fmt.Errorf("collection %s already exists", collName)
		}
	}

	coll := db.Collection(collName)
	mod := mongo.IndexModel{
		Keys: bson.M{
			// ascending order
			"user": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	_, err = coll.Indexes().CreateOne(ctx, mod)
	if err != nil {
		return fmt.Errorf("failed to create index: %v", err)
	}

	return nil
}
