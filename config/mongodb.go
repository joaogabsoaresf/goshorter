package config

import (
	"context"

	"github.com/joaogabsoaresf/goshorter/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func InitializeMongoDB() (*mongo.Collection, error) {
	logger := GetLogger("mongodb")
	mongoHost := "mongodb://localhost:27017/"
	dbName := "goshorter"
	collectionName := "url"

	clientOptions := options.Client().ApplyURI(mongoHost)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Errorf("mongodb connect has failed. error: %v", err)
		defer client.Disconnect(ctx)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Errorf("mongodb ping has failed. error: %v", err)
		return nil, err
	}

	dbExists, err := checkDatabaseExistence(ctx, client, dbName)
	if err != nil {
		logger.Errorf("error to check db existence: %v", err)
		return nil, err
	}

	if !dbExists {
		if err := createDatabaseAndCollection(ctx, client, dbName, collectionName); err != nil {
			return nil, err
		}
		logger.Info("success to create the database and the collection!")
	}

	db := client.Database(dbName).Collection(collectionName)

	return db, nil
}

func checkDatabaseExistence(ctx context.Context, client *mongo.Client, dbName string) (bool, error) {
	db, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		return false, err
	}

	for _, db := range db {
		if db == dbName {
			return true, nil
		}
	}

	return false, nil
}

func createDatabaseAndCollection(ctx context.Context, client *mongo.Client, dbName string, collectionName string) error {
	logger.Infof("database with name: %s not found, creating...", dbName)
	db := client.Database(dbName).Collection(collectionName)
	logger.Info("collection created! creating a empty document...")
	_, err := db.InsertOne(ctx, bson.M{})
	if err != nil {
		logger.Errorf("error to check db existence: %v", err)
		return err
	}

	return nil
}

func CreateUrlDocument(db *mongo.Collection, url *schemas.Url) error {
	_, err := db.InsertOne(ctx, url)
	if err != nil {
		return err
	}
	return nil
}

func FindDocumentFilter(db *mongo.Collection, filter bson.M) interface{} {
	var result interface{}
	info := db.FindOne(ctx, filter).Decode(&result)
	if info != nil {
		logger.Infof("no document find: %v", info)
		return nil
	}
	return result
}
