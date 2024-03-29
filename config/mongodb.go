package config

import (
	"context"
	"fmt"

	"github.com/joaogabsoaresf/goshorter/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func mongoHost(secrets *Secrets) string {
	if secrets.Env == "PROD" {
		mongoHostString := fmt.Sprintf("mongodb+srv://%s:%s@%s/", secrets.MongoDBUsername, secrets.MongoDBPassword, secrets.MongoDBHost)
		return mongoHostString
	} else {
		return secrets.LocalMongoDBHost
	}
}

func InitializeMongoDB(collectionNameOptional ...string) (*mongo.Database, error) {
	logger := GetLogger("mongodb")
	secrets := GetSecrets()
	mongoHost := mongoHost(secrets)
	dbName := "goshorter"

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
		var collectionName string
		collectionName = collectionNameOptional[0]
		if err := createDatabaseAndCollection(ctx, client, dbName, collectionName); err != nil {
			return nil, err
		}
		logger.Info("success to create the database and the collection!")
	}

	db := client.Database(dbName)

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

func CreateUrlDocument(db *mongo.Database, collectionName string, url *schemas.Url) error {
	_, err := db.Collection(collectionName).InsertOne(ctx, url)
	if err != nil {
		return err
	}
	return nil
}

func FindDocumentFilter(db *mongo.Database, collectionName string, filter bson.M) *schemas.UrlResponse {
	var result *schemas.UrlResponse
	info := db.Collection(collectionName).FindOne(ctx, filter).Decode(&result)
	if info != nil {
		logger.Infof("no document find: %v", info)
		return nil
	}
	return result
}
