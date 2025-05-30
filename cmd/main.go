package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shehbaazsk/go-commerce/config"
	"github.com/shehbaazsk/go-commerce/middlewares"
)

func main() {
	//Initialize configuration
	config.InitConfig()

	// Initialize Gin router
	router := gin.New()

	// Attach middleware
	router.Use(gin.Recovery())
	router.Use(middlewares.CustomLogger())

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "API is running",
		})
	})

	// Start the server
	log.Printf("Starting server on %s:%s", config.App.HOST, config.App.Port)
	router.Run(config.App.HOST + ":" + config.App.Port)

}
