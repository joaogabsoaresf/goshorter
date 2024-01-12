package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/goshorter/schemas"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUserHanlder(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.ValidateUserRequest(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	filterEmail := bson.M{"email": request.Email}
	if emailExist := schemas.FindUser(db, "user", filterEmail); emailExist != nil {
		logger.Errorf("email already exist in database: %v", request.Email)
		sendError(ctx, http.StatusConflict, "email already exist in database")
		return
	}

	filterSubDomain := bson.M{"subdomain": request.SubDomain}
	if subDomainExist := schemas.FindUser(db, "user", filterSubDomain); subDomainExist != nil {
		logger.Errorf("subdomain already exist in database: %v", request.SubDomain)
		sendError(ctx, http.StatusConflict, "subdomain already exist in database")
		return
	}

	user := schemas.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: request.Password,
		SubDomain:    request.SubDomain,
		IsAdmin:      false,
		UserID:       1,
	}

	if err := schemas.CreateNewUser(db, "user", &user); err != nil {
		logger.Errorf("user creation error: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "user", user)
}
