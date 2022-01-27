package aauth

import (
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSignup(t *testing.T) {
	client, ctx, close := ConnectToDB()
	defer close()

	DB_NAME = "aauth_test"
	err := Signup("hi", "ho")
	
	if err == errors.New("username already taken") {
		user := User{}

		client.Database(DB_NAME).Collection("users").FindOne(ctx, bson.D{
			primitive.E{ Key: "username", Value: "hi" },
		}).Decode(&user)

		if user.Username == "" {
			t.Error("Username already taken.")
		}
	} else if err == nil {
		user := User{}

		client.Database(DB_NAME).Collection("users").FindOne(ctx, bson.D{
			primitive.E{ Key: "username", Value: "hi" },
		}).Decode(&user)

		if user.Username == "" {
			t.Error("Signup did not add a user.")
		}
	}

	DB_NAME = "aauth"
}