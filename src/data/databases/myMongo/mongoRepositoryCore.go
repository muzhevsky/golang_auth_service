package myMongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-app/utils/errorsAndPanics"
)

type MongoRepositoryCore struct {
	database *mongo.Database
}

func (core *MongoRepositoryCore) InsertOne(collectionName string, content interface{}) {
	collection := core.database.Collection(collectionName)
	_, err := collection.InsertOne(context.Background(), content)
	errorsAndPanics.HandleError(err)
}

func (core *MongoRepositoryCore) InsertMany(collectionName string, content []interface{}) {
	collection := core.database.Collection(collectionName)
	_, err := collection.InsertMany(context.Background(), content)
	errorsAndPanics.HandleError(err)
}

func (core *MongoRepositoryCore) FindOne(collectionName string, filter interface{}) *mongo.SingleResult {
	collection := core.database.Collection(collectionName)
	return collection.FindOne(context.Background(), filter)
}

func (core *MongoRepositoryCore) Find(collectionName string, filter interface{}) *mongo.Cursor {
	collection := core.database.Collection(collectionName)
	result, err := collection.Find(context.Background(), filter)
	errorsAndPanics.HandleError(err)
	return result
}

func (core *MongoRepositoryCore) UpdateOne(collectionName string, filter interface{}, update interface{}) {
	collection := core.database.Collection(collectionName)
	_, err := collection.UpdateOne(context.Background(), filter, update)
	errorsAndPanics.HandleError(err)
}

func (core *MongoRepositoryCore) DeleteOne(collectionName string, filter interface{}) {
	collection := core.database.Collection(collectionName)
	_, err := collection.DeleteOne(context.Background(), filter)
	errorsAndPanics.HandleError(err)
}
