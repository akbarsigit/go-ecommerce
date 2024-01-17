package controllers

import (
	"context"
	"ecommerce/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invaid Search Index"})
			c.Abort()
			return
		}

		address, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(500, "internal server error")
		}
		
		var addresses models.Address

		addresses.Address_id = primitive.NewObjectID()
		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		defer cancel()

		match_filter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value:bson.D{primitive.E{Key:"path", Value:"$address"}}}}
		group := bson.D{{Key: "$group", Value:bson.D{primitive.E{Key:"_id", Value:"$address_id"}, {Key:"count", Value:bson.D{primitive.E{Key:"$sum", Value: 1}}}}}}

		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter, unwind, group})
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}
		
	}
}

func EditHomeAddress() gin.HandlerFunc {}

func EditWorkAddress() gin.HandlerFunc {}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invaid Search Index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		user_id_hex, err := primitive.ObjectIDFromHex(user_id)

		if err != nil {
			c.IndentedJSON(500, "internal server error")
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: user_id_hex}}
		update := bson.D{{Key:"$set", Value: bson.D{primitive.E{Key:"address", Value:addresses}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "Cannot Delete the Address")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Success Delete Address")
	}
}
