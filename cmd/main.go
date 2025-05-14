package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shehbaazsk/go-commerce/config"
	"github.com/shehbaazsk/go-commerce/middlewares"
)

func main() {

	// Load environment variables and config
	config.LoadConfig()

	// Initialize Gin router
	router := gin.Default()

	// Apply custom logger middleware
	router.Use(middlewares.CustomGinLogger())

	router.Use(gin.Recovery())

	// Health check route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Start the server
	log.Println("Starting server on port " + config.App.AppPort)
	if err := router.Run(":" + config.App.AppPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
