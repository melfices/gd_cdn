package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/melfices/go_cdn/controllers"
)

func CDNRoute(router *gin.Engine) {
	router.POST("/getcdn", controllers.GetCDN())
}
