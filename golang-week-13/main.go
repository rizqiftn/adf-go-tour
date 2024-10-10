package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type studentMongo struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connectMongo() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

func insert() {
	db, err := connectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, studentMongo{"Wick", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, studentMongo{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

func find() {
	db, err := connectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]studentMongo, 0)
	for csr.Next(ctx) {
		var row studentMongo
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Name  :", result[0].Name)
		fmt.Println("Grade :", result[0].Grade)
	}
}
func update() {
	db, err := connectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "Wick"}
	var changes = studentMongo{"John Wick", 1}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update success!")
}
func remove() {
	db, err := connectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "John Wick"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove success!")
}

func aggregate() {
	db, err := connectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	pipeline := make([]bson.M, 0)
	err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
		[
			{ "$group": {
				"_id": null,
				"Total": { "$sum": 1 }
			} },
			{ "$project": {
				"Total": 1,
				"_id": 0
			} }
		]
	`)), true, &pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Total :", result[0]["Total"])
	}
}

func main() {
	// insert()
	// find()
	// update()
	// remove()
	aggregate()
}
