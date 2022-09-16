package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const appName = "NyashkaButlerBot"

func Init(ctx context.Context) *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/").SetAppName(appName)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Can't connect to database: %s", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Can't ping database: %s", err)
	}
	return client.Database(appName)
}
