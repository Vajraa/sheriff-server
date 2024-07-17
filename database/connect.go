package database

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func SetupMongoDB() (*mongo.Collection, *mongo.Client, context.Context, context.CancelFunc) {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error Loading Envs")
	}
	dbUrl := os.Getenv("MONGO_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		slog.Error("Mongo DB Connect issue", "error", err)
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		slog.Error("Mongo DB ping issue", "error", err)
		panic(err)
	}
	collection := client.Database("sheriff").Collection("Users")
	slog.Info("Database connected Successfully")
	return collection, client, ctx, cancel
}

func CloseConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(context); err != nil {
			panic(err)
		}
		slog.Info("MongoDB Connection Closed")
	}()
}
