package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Collections map[string]*mongo.Collection
type DB struct {
	db          *mongo.Database
	collections Collections
}

func New(ctx context.Context, appName string) *DB {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/").SetAppName(appName)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Panicf("Can't connect to database: %s", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Panicf("Can't ping database: %s", err)
	}
	db := &DB{
		db:          client.Database(appName),
		collections: make(Collections),
	}
	return db
}
