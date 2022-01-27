package aauth

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_NAME = "aauth"
var URI = ""

func ConfigureURI(uri string) {
	URI = uri
}

func ConnectToDB() (*mongo.Client, context.Context, func()) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	CheckError(err)
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	CheckError(err)
	return client, ctx, close
}


type User struct {
	Username string `json:"username"`
	Token string `json:"token"`
	UserID int64 `json:"user_id"`
}