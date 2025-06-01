package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/shehbaazsk/go-commerce/config"
	"github.com/shehbaazsk/go-commerce/middlewares"
)

func main() {
	// Initialize configuration
	config.InitConfig()

	// Connect to DB
	dbpool, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	fmt.Println("Connected to Database successfully!")

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
