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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var validate = validator.New()

var restaurantCollection *mongo.Collection = database.OpenCollection(database.Client, "restaurant")

// add a restaurant to db (admin side)
func AdminAddRestaurantToDB(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var restaurant models.Restaurant

	if err := c.BindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(restaurant)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	restaurant.ID = primitive.NewObjectID()

	result, insertErr := restaurantCollection.InsertOne(ctx, restaurant)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error adding restaurant": insertErr.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// get all restaurants
func GetRestaurants(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var restaurants []bson.M

	cursor, err := restaurantCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding restaurant collection": err.Error()})
		return
	}

	if err = cursor.All(ctx, &restaurants); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting restaurant cursor": err.Error()})
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

// // get all orders by the waiter's name
// func GetOrdersByWaiter(c *gin.Context) {

// 	waiter := c.Params.ByName("waiter")

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	var orders []bson.M

// 	cursor, err := orderCollection.Find(ctx, bson.M{"server": waiter})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	if err = cursor.All(ctx, &orders); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	defer cancel()

// 	fmt.Println(orders)

// 	c.JSON(http.StatusOK, orders)
// }

// get all food items by the restaurant's id
func GetFoodByRestaurantID(c *gin.Context) {
	//have restaurant id
	//get all food id
	//search food with food id
	//return food array
	restaurantID := c.Params.ByName("restaurant_id")
	docID, _ := primitive.ObjectIDFromHex(restaurantID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var food []bson.M

	opts := options.Find().SetSort(bson.D{{"_id", 1}})
	cursor, err := foodCollection.Find(ctx, bson.M{"restaurant_id": docID}, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding food by restaurant ID": err.Error()})
		return
	}

	if err = cursor.All(ctx, &food); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting food by restaurant ID cursor": err.Error()})
		return
	}

	c.JSON(http.StatusOK, food)
}

// // update a waiter's name for an order
// func UpdateWaiter(c *gin.Context) {

// 	orderID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(orderID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	type Waiter struct {
// 		Server *string `json:"server"`
// 	}

// 	var waiter Waiter

// 	if err := c.BindJSON(&waiter); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	result, err := orderCollection.UpdateOne(ctx, bson.M{"_id": docID},
// 		bson.D{
// 			{"$set", bson.D{{"server", waiter.Server}}},
// 		},
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	defer cancel()

// 	c.JSON(http.StatusOK, result.ModifiedCount)

// }

// // update the order
// func UpdateOrder(c *gin.Context) {

// 	orderID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(orderID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	var order models.Order

// 	if err := c.BindJSON(&order); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	validationErr := validate.Struct(order)
// 	if validationErr != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 		fmt.Println(validationErr)
// 		return
// 	}

// 	result, err := orderCollection.ReplaceOne(
// 		ctx,
// 		bson.M{"_id": docID},
// 		bson.M{
// 			"dish":   order.Dish,
// 			"price":  order.Price,
// 			"server": order.Server,
// 			"table":  order.Table,
// 		},
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	defer cancel()

// 	c.JSON(http.StatusOK, result.ModifiedCount)
// }

// // delete an order given the id
// func DeleteOrder(c *gin.Context) {

// 	orderID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(orderID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	result, err := orderCollection.DeleteOne(ctx, bson.M{"_id": docID})

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	defer cancel()

// 	c.JSON(http.StatusOK, result.DeletedCount)

// }
