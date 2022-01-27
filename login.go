package aauth

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(username string, password string) error {
	token := GenToken(username, password)
	client, ctx, close := ConnectToDB()
	user := User{}
	defer close()

	client.Database(DB_NAME).Collection("users").FindOne(ctx, bson.D{
		primitive.E{ Key: "username", Value: username },
	}).Decode(&user)

	if user.Username == "" {
		return errors.New("invalid user")
	} else {
		if user.Token != token {
			return errors.New("wrong password")
		}
	}

	return nil
}