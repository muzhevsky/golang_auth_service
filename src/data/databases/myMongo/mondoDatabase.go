package myMongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang-app/src/data/databases"
	"golang-app/utils/errorsAndPanics"
)

type database struct {
	config         *databases.DatabaseConfig
	client         *mongo.Client
	repositoryCore *MongoRepositoryCore
}

func NewMongoClient(config *databases.DatabaseConfig) *database {
	return &database{config, nil, nil}
}

func (db *database) RepositoryCore() *MongoRepositoryCore {
	return db.repositoryCore
}

func (db *database) Connect() error {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.config.ConnectionString()))
	errorsAndPanics.HandleError(err)

	err = mongoClient.Ping(context.TODO(), readpref.Primary())
	errorsAndPanics.HandleError(err)

	db.client = mongoClient
	db.repositoryCore = &MongoRepositoryCore{database: mongoClient.Database(db.config.DatabaseName())}
	return nil
}

func (db *database) Disconnect() error {
	err := db.client.Disconnect(context.TODO())
	return err
}
