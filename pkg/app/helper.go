package app

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func writeResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		panic(err)
	}
}

func GetMongoClient() *mongo.Client {
	uri := "mongodb://admin:Faheem%40321%23@localhost:27017/?maxPoolSize=20&w=majority"
	clientOptions := options.Client().ApplyURI(uri)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	collection := client.Database("userdb").Collection("User")
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}, {Key: "email", Value: 1}, {Key: "phone", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
