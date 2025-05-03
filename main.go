package main

import (
	"log"
	"os"
	"ticket-system/config"
	"ticket-system/entity"
	route "ticket-system/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	// Connect to database
	config.InitDB()

	// Auto migrate tables
	if err := config.DB.AutoMigrate(
		&entity.User{},
		&entity.Event{},
		&entity.Ticket{},
		&entity.RevenueSummary{},
		&entity.EventTicketReport{},
	); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	// Setup router
	router := gin.Default()
	route.ConnectRoutes(router)

	// Get port from env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
