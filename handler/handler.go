package handler

import (
	"github.com/joaogabsoaresf/goshorter/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	logger *config.Logger
	db     *mongo.Collection
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMongoDB()
}
