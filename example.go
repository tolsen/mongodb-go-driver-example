package mongodb_go_driver_example

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RunExample() {
	client, err := mongo.NewClient()

	coll := client.Database("sample_mflix").Collection("movies")
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", "The Room"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return
		}
		panic(err)
	}
}
