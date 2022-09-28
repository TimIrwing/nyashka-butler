package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

const appName = "NyashkaButlerBot"
const DefaultURI = "mongodb://localhost:27017/"

type DB struct {
	orig *mongo.Database
}

func Init(ctx context.Context, uri string) *DB {
	clientOptions := options.Client().ApplyURI(uri).SetAppName(appName)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Can't connect to database: %s", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Can't ping database: %s", err)
	}
	return &DB{orig: client.Database(appName)}
}

func (db DB) GetChatCollection(chatID int64) *Collection {
	name := strconv.FormatInt(chatID, 10)
	return &Collection{orig: db.orig.Collection(name), name: name}
}