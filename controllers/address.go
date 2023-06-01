package controllers

import (
	"E-Commerce/logger"
	"E-Commerce/models"
	"context"
	"log"
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
		log.Println(user_id)
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid Code"})
			c.Abort()
			return
		}

		address, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(500, "Invalid User ID")
			log.Println(err)
		}

		var addresses models.Address

		addresses.Address_ID = primitive.NewObjectID()

		if err = c.BindJSON(&addresses); err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(http.StatusBadRequest, "Invalid Body")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		match_filter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$address_id"}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}

		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter, unwind, group})
		if err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(500, "Internal Server Error")
		}

		var addressinfo []bson.M
		if err = pointcursor.All(ctx, &addressinfo); err != nil {
			logger.LogError(err, logger.GetFileName())
			panic(err)
		}

		var size int32
		for _, address_no := range addressinfo {
			count := address_no["count"]
			size = count.(int32)
		}
		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
			_, err := UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				logger.LogError(err, logger.GetFileName())
				log.Println(err)
			}

		} else {
			c.JSON(400, "Not Alowed")
		}

		defer cancel()
		ctx.Done()

		c.JSON(200, "Successfully added address")
	}

}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid Id"})
			c.Abort()
			return
		}

		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(500, "Internal Server Error")
		}

		var editaddress models.Address
		if err := c.BindJSON(&editaddress); err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(http.StatusBadRequest, "Invalid Body")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.0.house", Value: editaddress.House}, {Key: "address.0.street", Value: editaddress.Street}, {Key: "address.0.city", Value: editaddress.City}, {Key: "address.0.pincode", Value: editaddress.Pincode}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(500, "Something Went Wrong")
			return
		}

		defer cancel()
		ctx.Done()
		c.JSON(200, "Successfully updated address")
	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid Id"})
			c.Abort()
			return
		}

		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(500, "Internal Server Error")
		}

		var editaddress models.Address
		if err := c.BindJSON(&editaddress); err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(http.StatusBadRequest, "Invalid Body")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address.1.house", Value: editaddress.House}, {Key: "address.1.street", Value: editaddress.Street}, {Key: "address.1.city", Value: editaddress.City}, {Key: "address.1.pincode", Value: editaddress.Pincode}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(500, "Something Went Wrong")
			return
		}

		defer cancel()
		ctx.Done()
		c.JSON(200, "Successfully updated address")
	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid search index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		userobj_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(500, "Internal Server Error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: userobj_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			logger.LogError(err, logger.GetFileName())
			c.JSON(404, "Couldn't delete addresses")
			return
		}

		defer cancel()
		ctx.Done()
		c.JSON(200, "Successfully deleted addresses")
	}
}
