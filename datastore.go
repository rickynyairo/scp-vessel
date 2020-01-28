package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// CreateClient -
func CreateClient(uri string) (*mongo.Client, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	cancelFunc()
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}
