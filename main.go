package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
)

type BankAccount struct {
	Name    string  `json:"name"`
	GovID   string  `json:"govid"`
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

var (
	mongoURI    = "mongodb://localhost:27017"
	database    = "bank"
	collection  = "accounts"
	mongoClient *mongo.Client
)

func initMongoDB() error {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB!")
	mongoClient = client
	return nil
}

func main() {
	err := initMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	app := gofr.New()

	// POST
	app.POST("/account/create", func(ctx *gofr.Context) (interface{}, error) {
		var account BankAccount
		err := json.NewDecoder(ctx.Request().Body).Decode(&account)
		if err != nil {
			return nil, err
		}
		accountsCollection := mongoClient.Database(database).Collection(collection)
		_, err = accountsCollection.InsertOne(context.Background(), account)
		if err != nil {
			return nil, err
		}
		return "Bank account created successfully", nil
	})

	// GET
	app.GET("/accounts/list", func(ctx *gofr.Context) (interface{}, error) {
		var accounts []BankAccount
		accountsCollection := mongoClient.Database(database).Collection(collection)
		cursor, err := accountsCollection.Find(context.Background(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.Background())
		if err := cursor.All(context.Background(), &accounts); err != nil {
			return nil, err
		}
		return accounts, nil
	})

	// PUT
	app.PUT("/account/update/{id}", func(ctx *gofr.Context) (interface{}, error) {
		urlPath := ctx.Request().URL.Path
		path := strings.Split(urlPath, "/")
		govID := path[3]
		accountID, err := primitive.ObjectIDFromHex(govID)
		if err != nil {
			return nil, err
		}
		var updateParams map[string]interface{}
		err = json.NewDecoder(ctx.Request().Body).Decode(&updateParams)
		if err != nil {
			return nil, err
		}
		accountsCollection := mongoClient.Database(database).Collection(collection)
		filter := bson.M{"_id": accountID}
		update := bson.M{"$set": updateParams}
		_, err = accountsCollection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return "Bank account information updated successfully", nil
	})

	// DELETE
	app.DELETE("/account/remove/{id}", func(ctx *gofr.Context) (interface{}, error) {
		urlPath := ctx.Request().URL.Path
		path := strings.Split(urlPath, "/")
		govID := path[3]
		accountID, err := primitive.ObjectIDFromHex(govID)
		accountsCollection := mongoClient.Database(database).Collection(collection)
		_, err = accountsCollection.DeleteOne(context.Background(), bson.M{"_id": accountID})
		if err != nil {
			return nil, err
		}
		return "Bank account removed successfully", nil
	})
	app.Start()
}
