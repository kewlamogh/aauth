package aauth

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup(username string, password string) error {
	client, ctx, close := ConnectToDB()
	user := User{}
	defer close()

	client.Database(DB_NAME).Collection("users").FindOne(ctx, bson.D{
		primitive.E{Key: "username", Value: username},
	}).Decode(&user)

	if user.Username == "" {
		userID, err := client.Database(DB_NAME).Collection("users").CountDocuments(ctx, bson.D{})
		CheckError(err)

		client.Database(DB_NAME).Collection("users").InsertOne(ctx, bson.D{
			primitive.E{Key: "username", Value: username},
			primitive.E{Key: "token", Value: GenToken(username, password)},
			primitive.E{Key: "user_id", Value: userID},
		})

		return nil
	}

	return errors.New("username already taken")
}
