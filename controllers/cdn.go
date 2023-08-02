package controllers

import (
	//other import goes here

	"go_cdn/configs" //add this
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

var cdnCollection *mongo.Collection = configs.GetCollection(configs.DB, configs.EnvCollection())

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
