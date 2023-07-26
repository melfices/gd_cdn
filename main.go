package main

import (
	"github.com/gin-gonic/gin"
	"github.com/melfices/go_cdn/configs"
	"github.com/melfices/go_cdn/routes"
)

func main() {
	router := gin.Default()
	configs.ConnectDB()
	routes.CDNRoute(router)
	router.Run("localhost:3000")
}
