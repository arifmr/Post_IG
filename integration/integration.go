package integration

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongodb() (*mongo.Client, error) {
	post, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = post.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return post, nil
}
