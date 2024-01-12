package handler

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/goshorter/config"
	"github.com/joaogabsoaresf/goshorter/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUrlHandler(ctx *gin.Context) {
	request := CreateUrlRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	url := schemas.Url{
		OriginalPath: request.OriginalPath,
		ShorterID:    createNonExistentID(),
	}

	if err := config.CreateUrlDocument(db, "url", &url); err != nil {
		logger.Errorf("url creation error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create", url)
}

func createNonExistentID() string {
	id := createNewUniqueID(5)
	if idExist(id) {
		for !idExist(id) {
			id = createNewUniqueID(5)
			if !idExist(id) {
				break
			}
		}
	}
	return id
}

func createNewUniqueID(lenght int) string {

	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	result := ""

	for i := 0; i < lenght; i++ {
		randIndex := rand.Intn(len(characters))
		result += string(characters[randIndex])
	}

	return result
}

func idExist(id string) bool {
	filter := bson.M{"shorterid": id}
	exist := config.FindDocumentFilter(db, "url", filter)
	return exist != nil
}
