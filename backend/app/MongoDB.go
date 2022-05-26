package app

import (
	"context"
	"fmt"
	"go-fiber-app/helper"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conn *mongo.Client

type MongoDB struct{}

func (m MongoDB) NewDB() interface{} {

	if conn != nil {
		return conn
	}

	fmt.Println("Connecting Database...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	_ = cancel

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		Username:      os.Getenv("DB_USERNAME"), // your mongodb user
		Password:      os.Getenv("DB_PASSWORD"), // ...and mongodb
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_HOST")).SetAuth(credential)

	client, err := mongo.NewClient(clientOptions)
	helper.PanicIfNeeded(err)
	// connect to mongo
	err = client.Connect(ctx)
	helper.PanicIfNeeded(err)
	// test connection to mongo
	err = client.Ping(ctx, nil)
	helper.PanicIfNeeded(err)
	fmt.Println("Connected")

	conn = client

	return conn
}
