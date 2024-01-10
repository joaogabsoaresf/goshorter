package config

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db      *mongo.Collection
	logger  *Logger
	secrets *Secrets
)

func Init() error {
	var err error
	secrets = InitializeSecrets()

	db, err = InitializeMongoDB()
	if err != nil {
		return fmt.Errorf("erro in mongodb init: %v", err)
	}

	return nil
}

func GetMongoDB() *mongo.Collection {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetSecrets() *Secrets {
	return secrets
}
