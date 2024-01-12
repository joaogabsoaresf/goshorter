package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/goshorter/config"
	"go.mongodb.org/mongo-driver/bson"
)

func ShowUrlHandler(ctx *gin.Context) {
	shorter_id := ctx.Query("shorter_id")
	if shorter_id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("shorter_id", "queryParameter").Error())
		return
	}

	filter := bson.M{"shorterid": shorter_id}
	result := config.FindDocumentFilter(db, "url", filter)
	if result == nil {
		logger.Errorf("error url not found")
		sendError(ctx, http.StatusNotFound, errDocumentNotFound("shorter_id").Error())
		return
	}

	sendSuccess(ctx, "find", result)
}
