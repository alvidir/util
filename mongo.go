package util

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConn(uri string, timeout time.Duration) (client *mongo.Client, cancel context.CancelFunc, err error) {
	var ctx context.Context
	ctx, cancel = context.WithTimeout(context.Background(), timeout)

	options := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(ctx, options)
	return
}
