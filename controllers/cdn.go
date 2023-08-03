package controllers

import (
	//other import goes here

	"context"
	"fmt"
	"go_cdn/configs" //add this
	"go_cdn/models"
	"go_cdn/responses"
	"go_cdn/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var cdnCollection *mongo.Collection = configs.GetCollection(configs.DB, configs.EnvCollection())

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func GetCDNDetail(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cdn_param := new(models.CDNParam)
	defer cancel()
	if err := c.Bind(cdn_param); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(cdn_param)
	var cdn models.CDN
	filter := bson.D{{Key: "cdn", Value: cdn_param.CDN}}
	err := cdnCollection.FindOne(ctx, filter).Decode(&cdn)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": cdn}})
}
func GetCDN(c echo.Context) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	stype := new(models.SysType)
	// defer cancel()
	if err := c.Bind(stype); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	d := time.Now()
	twodigityear := d.Year() % 1e2
	dstr := fmt.Sprintf("%02d%02d%02d", twodigityear, d.Month(), d.Day())
	rdigit, _ := utils.GenRanDigit()
	cdn := dstr + rdigit
	chkdigit := strconv.Itoa(utils.CheckDigit(cdn))
	cdn = cdn + chkdigit
	fmt.Println(cdn)
	// var cdn models.CDN
	// filter := bson.D{{"cdn", cdn_param.CDN}}
	// err := cdnCollection.FindOne(ctx, filter).Decode(&cdn)

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	// }
	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": cdn}})
}
