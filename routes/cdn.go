package routes

import (
	"go_cdn/controllers"

	"github.com/labstack/echo/v4"
)

func CDNRoute(e *echo.Echo) {
	e.POST("/getcdn", controllers.Hello)
}
