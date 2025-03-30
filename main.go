package main

import (
	"backend/config"
	"backend/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	if err := config.ConnectDB(); err != nil { // Now correctly checking the error
		log.Fatal("❌ Failed to connect to database. Exiting...")
	}

	// Create a new Gin router
	router := gin.Default()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins (change to frontend URL in production)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Register authentication routes
	routes.SetupAuthRoutes(router)

	// Start the server on port 8080
	log.Println("🚀 Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}
