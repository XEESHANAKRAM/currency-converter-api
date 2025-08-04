// main.go - Entry point of our application
package main

import (
	"log"
	"os"

	"currency-converter/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router with default middleware
	r := gin.Default()

	// Configure CORS to allow frontend access
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	// Serve static files (HTML, CSS, JS)
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")

	// API routes
	api := r.Group("/api/v1")
	{
		api.GET("/convert", handlers.ConvertCurrency)
		api.GET("/rates", handlers.GetExchangeRates)
		api.GET("/health", handlers.HealthCheck)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
