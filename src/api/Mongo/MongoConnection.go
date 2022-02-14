package Mongo

import (
	"Packages/src/Configs"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
)

var once sync.Once

var Collection *mongo.Collection

func GetMongoSingletonCollection(client *mongo.Client, configs *Configs.AppConfig) *mongo.Collection {
	once.Do(func() {
		collection, _ := GetMongoDbCollection(client, configs)
		Collection = collection
	})
	return Collection
}

func GetMongoClient(configs *Configs.AppConfig) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), configs.MongoConnectionDuration)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.MongoConnectionURI))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetMongoDbCollection(client *mongo.Client, configs *Configs.AppConfig) (*mongo.Collection, error) {
	collection := client.Database(configs.DBName).Collection(configs.CollectionName)
	return collection, nil
}
