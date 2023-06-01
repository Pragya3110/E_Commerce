package database

import (
	"E-Commerce/logger"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {

	err := godotenv.Load(".env")
	if err != nil {
		logger.LogError(err, logger.GetFileName())
		panic(err)
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("URI")))

	if err != nil {
		logger.LogError(err, logger.GetFileName())
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logger.LogError(err, logger.GetFileName())
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.LogError(err, logger.GetFileName())
		return nil
	}

	fmt.Println("Successfully connected to mongodb")
	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(os.Getenv("DB_NAME")).Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productcollection *mongo.Collection = client.Database(os.Getenv("DB_NAME")).Collection(collectionName)
	return productcollection
}
