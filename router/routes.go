package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaogabsoaresf/goshorter/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/url", handler.ShowUrlHandler)
		v1.POST("/url", handler.CreateUrlHandler)
		v1.POST("/register", handler.CreateUserHanlder)
	}

	router.GET("/:shortlink", handler.RedirectUrlHandler)
}
