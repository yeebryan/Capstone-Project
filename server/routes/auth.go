package routes

import (
	"context"
	"log"
	controller "server/controllers"
	helper "server/helpers"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("/signup", controller.SignUp)
	incomingRoutes.POST("/login", controller.Login)
}

// SignUp (Admin use)
func AdminSignUp(user models.User) models.User {
	validationErr := validate.Struct(user)
	if validationErr != nil {
		log.Panic(validationErr)
	}

	password, err := controller.HashPassword(*user.Password)
	if err != nil {
		log.Panic(err)
	}
	user.Password = &password
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()
	token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, user.UserID)
	user.Token = &token
	user.RefreshToken = &refreshToken

	_, insertErr := userCollection.InsertOne(context.Background(), user)
	if insertErr != nil {
		log.Panic(err)
	}

	return user
}
