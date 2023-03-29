package routes

import (
	"context"
	"net/http"
	"server/database"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cartCollection *mongo.Collection = database.OpenCollection(database.Client, "cart")

// get a cart by user ID
func GetCartByUserID(c *gin.Context) {
	userID := c.Params.ByName("user_id")
	docID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding user ID": err.Error()})
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var cart models.Cart

	err = cartCollection.FindOne(ctx, bson.M{"user_id": docID}).Decode(&cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding cart by user ID": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}
