package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Film struct {
	Id string `bson:"id"`
	Title string `bson:"title"`
}

func main() {
	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " + "See: " + docs + "usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		err := client.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()

	title := "Back to the Future"
	coll := client.Database("sample_mflix").Collection("movies")
	/*
	writeResult, err := coll.InsertOne(context.TODO(), Film{"id1", title})
	if err != nil {
		panic(err)
	}
	writeJsonData, err := json.MarshalIndent(writeResult, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", writeJsonData)
	fmt.Println("Write result:", writeResult)
	*/
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: "title", Value: title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}