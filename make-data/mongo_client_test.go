package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestMongoCli(t *testing.T) {
	properties := MongoProperties{
		Uri:      "mongodb://mongodb:mongodb@localhost:27017/admin", // 本地
		DataBase: "inner-drugs-auto",
	}
	client, err := properties.NewMongoClient()
	defer client.Disconnect(context.TODO())
	if err != nil {
		t.Fatal("connect fail")
	}

	// collection := client.Database(properties.DateBase).Collection("dc_standard_field_from_excel")
	collection := client.Database(properties.DataBase).Collection("dc_standard_field" + "_" + time.Now().String())
	// 查全部
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		t.Fatal("Find fail")
	}
	maps := make([]any, 1)
	err = cur.All(context.Background(), &maps)
	if err != nil {
		t.Fatal("Find err")
	}
}
