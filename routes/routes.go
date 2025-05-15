package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shehbaazsk/go-commerce/config"
)

func InitRoutes(r *gin.Engine) {

	// API Version group
	apiV1 := r.Group("/api/v1")
	{
		// Healthcheck route
		apiV1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})
		RegisterUserRoutes(apiV1, config.DB)

	}
}
