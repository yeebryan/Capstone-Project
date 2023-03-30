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

	fmt.Println("categoryPLS", category)
	fmt.Println("food typePLS", foodType)

	var query bson.M
	if category != "" && foodType != "" {
		query = bson.M{"category": bson.M{"$regex": category, "$options": "i"}, "foodtype": bson.M{"$regex": foodType, "$options": "i"}}
	} else if category != "" {
		query = bson.M{"category": bson.M{"$regex": category, "$options": "i"}}
	} else if foodType != "" {
		query = bson.M{"foodtype": bson.M{"$regex": foodType, "$options": "i"}}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing category or foodType parameter"})
		return
	}

	fmt.Println("Query yaya:", query)

	var foods []models.Food
	cur, err := foodCollection.Find(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch food", "details": err.Error()})
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

	if err := cur.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error", "details": err.Error()})
		return
	}

	if err := cur.Close(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close cursor", "details": err.Error()})
		return
	}

	fmt.Println("Foods FREAKYU:", foods)

	if len(foods) == 0 {
		c.JSON(http.StatusOK, gin.H{"foods": []models.Food{}})
		return
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(foods))
	food := foods[randomIndex]

	fmt.Println("Food WULALA items:", food)

	c.JSON(http.StatusOK, gin.H{"foods": []models.Food{food}})
}
