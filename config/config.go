package config

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db     *mongo.Collection
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeMongoDB()
	if err != nil {
		return fmt.Errorf("erro in mongodb init: %v", err)
	}

	return nil
}

func GetSQLite() *mongo.Collection {
	return db
}

func GetLogger(p string) *Logger {
	// init logger
	logger = NewLogger(p)
	return logger
}
