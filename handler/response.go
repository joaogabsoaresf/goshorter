package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/goshorter/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("url %s success", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message    string `json:"message"`
	ErrorCorde string `json:"errorCord"`
}

type CreateResponse struct {
	Message string              `json:"message"`
	Data    schemas.UrlResponse `json:"data"`
}
