// server.go
package main

import (
	"log"
	"os"

	"Server/db"
	"Server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
   // Initialize Firestore
   if err := db.InitFirestore(); err != nil {
      log.Fatalf("Error initializing Firestore: %v", err)
   }

   // Set up Gin router
   router := gin.Default()

   // Set up routes
   routes.SetupUserRoutes(router)
   routes.SetupNGORoutes(router)

   // Run the server
   port := os.Getenv("PORT")
   if port == "" {
      port = "8080"
   }
   router.Run(":" + port)
}
