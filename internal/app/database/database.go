package database

import (
	"context"
	"fmt"
	"github.com/DosyaKitarov/market-sniper/internal/pkg/env"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb() mongo.Client {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	name := env.GetEnvVariable("DB_USER_NAME")
	password := env.GetEnvVariable("DB_PASSWORD")
	uri := fmt.Sprintf("mongodb+srv://%s:%s@amazonscrapper.8fdyq9u.mongodb.net/?retryWrites=true&w=majority&appName=AmazonScrapper", name, password)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("AmazonScrapper").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return *client
}
