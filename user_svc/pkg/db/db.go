package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

func ConnectMongo() {
	client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatal(err, " - new database client failed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err, " connect database failed")
	}
	fmt.Println("connect mongo DB successfuly")
	defer cancel()
}

func GetDb() *mongo.Database {
	if err != nil {
		ConnectMongo()
	}
	return client.Database("users")
}

func DisconnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Disconnect(ctx)
	fmt.Println("MongoDb disconnected", ctx.Err())
}
