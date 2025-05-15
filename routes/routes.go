package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shehbaazsk/go-commerce/middlewares"
)

func InitRoutes(r *gin.Engine) {

	// API Version group
	apiV1 := r.Group("/api/v1")
	{
		public := apiV1.Group("/")
		// Healthcheck route
		public.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		private := apiV1.Group(("/"))
		private.Use(middlewares.AuthMiddleware())
	}
}
