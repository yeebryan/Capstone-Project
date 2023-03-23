package routes

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"server/database"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

// add a food item to db (admin side)
func AdminAddFoodToDB(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var food models.Food

	if err := c.BindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(food)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	food.ID = primitive.NewObjectID()

	result, insertErr := foodCollection.InsertOne(ctx, food)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error adding food": insertErr.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// get all food item (admin side)
func AdminGetFood(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var food []bson.M

	cursor, err := foodCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding food collection": err.Error()})
		return
	}

	if err = cursor.All(ctx, &food); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting food cursor": err.Error()})
		return
	}

	c.JSON(http.StatusOK, food)
}

func GetFoodByCategory(c *gin.Context) {
	category := c.Params.ByName("category")
	//add validation
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	food := []models.Food{}
	cursor, err := foodCollection.Find(ctx, bson.M{"category": bson.M{"$regex": category, "$options": "i"}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding food collection": err.Error()})
		return
	}

	if err = cursor.All(ctx, &food); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting food cursor": err.Error()})
		return
	}
	c.JSON(http.StatusOK, food)
}

// POST request FOOD
// generate food playlist
func FetchRandomFood(c *gin.Context) {
	category := c.Query("category")
	foodType := c.Query("foodType")

	var query bson.M
	if category != "" && foodType != "" {
		query = bson.M{"category": category, "food_type": foodType}
	} else if category != "" {
		query = bson.M{"category": category}
	} else if foodType != "" {
		query = bson.M{"food_type": foodType}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing category or foodType parameter"})
		return
	}

	fmt.Println("Query yaya:", query)

	var foods []models.Food
	cur, err := foodCollection.Find(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch food"})
		return
	}

	for cur.Next(context.Background()) {
		var food models.Food
		err := cur.Decode(&food)
		if err != nil {
			fmt.Println("Failed to decode food:", err)
			continue
		}
		foods = append(foods, food)
	}

	fmt.Println("Foods:", foods) // add this line to log the food items being returned

	if len(foods) == 0 {
		c.JSON(http.StatusOK, gin.H{"foods": []models.Food{}})
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(foods))
	food := foods[randomIndex]

	fmt.Println("Food lala items:", food)

	c.JSON(http.StatusOK, gin.H{"foods": []models.Food{food}})
}

// func GetFoodByRestaurantID(c *gin.Context) {

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

// // get all food items by the restaurant's id
// func GetFoodByRestaurantID(c *gin.Context) {

// 	restaurantID := c.Params.ByName("restaurant_id")
// 	docID, _ := primitive.ObjectIDFromHex(restaurantID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	var food []bson.M

// 	opts := options.Find().SetSort(bson.D{{"_id", 1}})
// 	cursor, err := foodCollection.Find(ctx, bson.M{"restaurant_id": docID}, opts)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error finding food by restaurant ID": err.Error()})
// 		return
// 	}

// 	if err = cursor.All(ctx, &food); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error getting food by restaurant ID cursor": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, food)
// }

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
