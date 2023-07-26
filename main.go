package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/melfices/go_cdn/configs"
	"github.com/melfices/go_cdn/routes"
)

func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	router := gin.Default()
	configs.ConnectDB()
	routes.CDNRoute(router)
	router.Run("localhost:3000")
}
