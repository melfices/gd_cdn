package controllers

import (
	//other import goes here

	"go_cdn/configs" //add this
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

var cdnCollection *mongo.Collection = configs.GetCollection(configs.DB, os.Getenv("COLLECTION"))

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
