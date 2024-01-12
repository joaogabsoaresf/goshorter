package router

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// gin.SetMode(gin.ReleaseMode) TO-DO Active if env == PROD
	r := gin.Default()

	r.Use(getAllowedSubdomains())

	initializeRoutes(r)

	r.Run(":8080")
}

func getAllowedSubdomains() gin.HandlerFunc {
	allowedSubdomains := []string{"admin", "lolis", "www"}

	return func(c *gin.Context) {
		host := c.Request.Host
		validSubdomain := false
		for _, subdomain := range allowedSubdomains {
			if strings.HasPrefix(host, subdomain+".") {
				validSubdomain = true
				break
			}
		}

		if !validSubdomain {
			c.JSON(403, gin.H{
				"error": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
