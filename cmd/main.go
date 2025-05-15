package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shehbaazsk/go-commerce/config"
	"github.com/shehbaazsk/go-commerce/middlewares"
	"github.com/shehbaazsk/go-commerce/routes"
)

func main() {

	// Load environment variables and config
	config.LoadConfig()

	// Initialize Gin router
	router := gin.Default()

	// Apply custom logger middleware
	router.Use(middlewares.CustomGinLogger())

	router.Use(gin.Recovery())

	routes.InitRoutes(router)

	// Start the server
	log.Println("Starting server on port " + config.App.AppPort)
	if err := router.Run(":" + config.App.AppPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
