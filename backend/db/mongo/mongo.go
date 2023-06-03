package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoConn(ctx context.Context, uri string) (*mongo.Client, error) {
	if len(uri) == 0 {
		logrus.Fatal("uri settings corrupt!")
		return nil, errors.New("传入无效uri")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetTimeout(time.Second*5).SetRetryReads(true))
	if err != nil {
		logrus.Error("Build connect failed")
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		logrus.Error("Connect failed", err)
		return nil, err
	} else {
		fmt.Println("Pong!")
	}
	return client, nil
}
