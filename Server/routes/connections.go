package routes

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create a mongodb client interactive instance :

func DBinstance() *mongo.Client {
	// get mongo url from .env file :
	err := godotenv.Load(".env")
	// check error state :
	if err != nil {
		log.Fatal("Error loading Env file :( ")
	}
	MongoDb := os.Getenv("MONGODB_URL")
	// try to make connexion to mongo atlas :
	Client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	// check connexion state :
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" Connected to MongoDb ")
	return Client
}

var Client *mongo.Client = DBinstance()

// create a function to getback my mongodb collection handler :
func OpenCollection(client *mongo.Client, collectionName string, dbName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(dbName).Collection(collectionName)
	return collection
}
