package routes

import (
	"context"
	"fmt"
	"net/http"
	"server/database"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// add a user to db (admin side)
func AdminAddUserToDB(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	user.ID = primitive.NewObjectID()

	result, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error adding user": insertErr.Error()})
		fmt.Println(insertErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// add multiple users to db (admin side)
func AdminAddMultipleUserToDB(c *gin.Context, users []models.User) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.BindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(users)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	for _, user := range users {
		user.ID = primitive.NewObjectID()
		_, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error adding multiple user": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, nil)
}

// get all users (admin side)
func AdminGetUsers(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var users []bson.M

	cursor, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding user collection": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting user cursor": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, users)
}
