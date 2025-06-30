package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBName = "db2025"
var MahasiswaCollection = "data_mahasiswa"
var UserCollection = "user"
var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect error: %v\n", err)
	}
	return client.Database(dbname)
}
