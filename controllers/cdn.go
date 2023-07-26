package controllers

import (
	//other import goes here
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/melfices/go_cdn/configs"
	"github.com/melfices/go_cdn/models"
	"github.com/melfices/go_cdn/responses"
	"go.mongodb.org/mongo-driver/bson" //add this
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cdnCollection *mongo.Collection = configs.GetCollection(configs.DB, "tb_cdn_pool")

func GetCDN() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.CDN
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := cdnCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}
