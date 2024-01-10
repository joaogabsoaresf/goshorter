package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// gin.SetMode(gin.ReleaseMode) TO-DO Active if env == PROD
	r := gin.Default()

	initializeRoutes(r)

	r.Run(":8080")
}
