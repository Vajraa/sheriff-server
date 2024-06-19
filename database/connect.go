package database

import (
	"context"
	"fmt"
	"log"
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
        log.Fatalf("Error loading .env file")
    }
	dbUrl := os.Getenv("MONGO_URL")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
	 panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
	 panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}
	collection := client.Database("sheriff").Collection("Users")
	return collection, client, ctx, cancel
   }
   
   func CloseConnection(client *mongo.Client, context context.Context, cancel context.CancelFunc) {
	defer func() {
	 cancel()
	 if err := client.Disconnect(context); err != nil {
	  panic(err)
	 }
	 fmt.Println("Close connection is called")
	}()
   }