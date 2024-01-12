package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/goshorter/config"
	"go.mongodb.org/mongo-driver/bson"
)

func RedirectUrlHandler(ctx *gin.Context) {
	shorter_id := ctx.Param("shortlink")
	if shorter_id == "" {
		sendError(ctx, http.StatusNotFound, errParamIsRequired("shorter_id", "Parameter").Error())
		return
	}

	filter := bson.M{"shorterid": shorter_id}
	result := config.FindDocumentFilter(db, "url", filter)
	if result == nil {
		logger.Errorf("error url not found")
		sendError(ctx, http.StatusNotFound, errDocumentNotFound("shorter_id").Error())
		return
	}

	sendRedirect(ctx, result.OriginalPath)
}
