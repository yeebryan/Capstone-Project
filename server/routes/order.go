package routes

import (
	"context"
	"log"
	"net/http"
	"server/database"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

// get all orders
func GetOrders(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var orders []bson.M

	cursor, err := orderCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding order collection": err.Error()})
		return
	}

	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting order cursor": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// get all order items by the user's id
func GetOrderCurrentUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userID := c.Value("uid")
	log.Println("userid: ", userID)
	userOID, err := primitive.ObjectIDFromHex(userID.(string))
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting user OID": err.Error()})
		return
	}
	log.Println("useroid: ", userOID)
	var result []bson.M

	pipeline := []bson.M{
		// bson.M{
		{"$match": bson.M{
			"user_id": userOID},
		},
	}
	log.Println(pipeline)
	cursor, err := orderCollection.Aggregate(ctx, pipeline)

	log.Println(cursor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding orders by user ID": err.Error()})
		return
	}

	if err = cursor.All(ctx, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting order cursor": err.Error()})
		return
	}
	log.Println("went here", result)
	c.JSON(http.StatusOK, result)
}
