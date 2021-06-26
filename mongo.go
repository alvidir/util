package util

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ErrCtxHasNoDeadline = "context has no deadline"
)

func MongoConn(ctx context.Context, uri string, opts ...*options.ClientOptions) (client *mongo.Client, err error) {
	if _, ok := ctx.Deadline(); !ok {
		err = errors.New(ErrCtxHasNoDeadline)
		return
	} else if err = ctx.Err(); err != nil {
		return
	}

	options := options.Client().ApplyURI(uri)
	opts = append(opts, options)

	client, err = mongo.Connect(ctx, opts...)
	return
}
