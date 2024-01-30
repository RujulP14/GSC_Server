// server.go
package main

import (
	"log"
	"os"

	"Server/db"
	"Server/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize Firestore
	if err := db.InitFirestore(); err != nil {
		log.Fatalf("Error initializing Firestore: %v", err)
	}

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set up Gin router
	router := gin.Default()

	// Set up routes
	routes.SetupUserRoutes(router)
	routes.SetupNGORoutes(router)
	routes.SetupChatbotRoutes(router)

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
