package main

import (
	"fmt"
	"log"

	"github.com/shehbaazsk/go-commerce/config"
	"github.com/shehbaazsk/go-commerce/router"
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
	router := router.SetupRouter(dbpool)

	for _, route := range router.Routes() {
		fmt.Printf("%-6s %s\n", route.Method, route.Path)
	}

	// Start the server
	log.Printf("Starting server on %s:%s", config.App.HOST, config.App.Port)
	router.Run(config.App.HOST + ":" + config.App.Port)

}
