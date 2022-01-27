package aauth

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestConnectFunction(t *testing.T) {
	client, ctx, close := ConnectToDB()
	obj := bson.D{ primitive.E{ Key: "key", Value: "value" } }
	v := struct{ Key string `json:"key"` }{}  

	defer close()

	DB_NAME = "aauth_test"
	
	client.Database(DB_NAME).Collection("samp").DeleteMany(ctx, obj)
	client.Database(DB_NAME).Collection("samp").InsertOne(ctx, obj)
	client.Database(DB_NAME).Collection("samp").FindOne(ctx, obj).Decode(&v)
	
	if v.Key != "key" {
		t.Errorf("Want %s, got %s", "key", v.Key)
	}

	DB_NAME = "aauth"
}