package main

import (
	"log"

	"go_cdn/configs"
	"go_cdn/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

//	func main() {
//		router := gin.Default()
//		configs.ConnectDB()
//		routes.CDNRoute(router)
//		router.Run("localhost:3000")
//	}
func main() {
	e := echo.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.CDNRoute(e)

	e.Logger.Fatal(e.Start(":3000"))
}
