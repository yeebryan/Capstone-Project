package routes

import (
	"context"
	"log"
	"net/http"
	"server/database"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetRestaurantByCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	category := c.Query("category")
	//add validation

	restaurants := []models.Restaurant{}
	cursor, err := restaurantCollection.Find(ctx, bson.M{"category": bson.M{"$regex": category, "$options": "i"}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding food collection": err.Error()})
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
	restaurantID := c.Params.ByName("restaurant_id")
	docID, err := primitive.ObjectIDFromHex(restaurantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting restaurant ID": err.Error()})
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var restaurant models.Restaurant

	err = restaurantCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&restaurant)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding food by restaurant ID": err.Error()})
		return
	}

	var food = make([]models.Food, len(restaurant.Menu.Menu))
	for i, foodID := range restaurant.Menu.Menu {
		err = foodCollection.FindOne(ctx, bson.M{"_id": foodID}).Decode(&food[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error finding food": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, food)
}

// add food item to cart
func AddFoodItemToCart(c *gin.Context) {
	//get userid from claims??
	//get foodid
	//get cart using userid + state != Completed
	//if found
	//	TODO:check if createdAt time more than 24h
	//		if yes
	//			clear cart
	//	see if food item exists
	//		yes = quantity + 1
	//	else add to array
	//if not found upsert new cart??
	//status ok,nil
	// userID := c.Value("uid")
	userID := c.Value("uid")
	userOID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting user OID": err.Error()})
		return
	}
	foodID := c.Params.ByName("food_id")
	foodDocID, err := primitive.ObjectIDFromHex(foodID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting food ID": err.Error()})
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var cart models.Cart

	err = cartCollection.FindOne(ctx, bson.M{"user_id": userOID, "state": models.StateInProcess}).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		//upsert cart
		log.Println("found nothing")
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding cart while adding food item": err.Error()})
		return
	}
	//found cart
	log.Println("cart before:", cart)
	var found bool
	for i, food := range cart.Items {
		log.Println(i, food)
		if food.ID == foodDocID {
			cart.Items[i].Quantity += 1
			cart.TotalPrice += food.Price
			found = true
			log.Println("adding")
			break
		}
	}
	//food not found in cart
	//get food info
	//append to cart.items
	if !found {
		var tempFood models.FoodItems
		err = foodCollection.FindOne(ctx, bson.M{"_id": foodDocID}).Decode(&tempFood)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error finding food item": err.Error()})
			return
		}
		tempFood.ID = foodDocID
		tempFood.Quantity = 1
		log.Println("Tempfood:", tempFood)
		cart.Items = append(cart.Items, tempFood)
	}

	log.Println("cart after:", cart) //should change quantity and cart total price
	update := bson.M{"$set": cart}
	_, err = cartCollection.UpdateByID(ctx, cart.ID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error updating cart while adding food item": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
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
