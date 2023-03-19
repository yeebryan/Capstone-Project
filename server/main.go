package main

import (
	"os"
	"server/middleware"
	"server/routes"

	// "server/testdata"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001"}
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "token"}
	router.Use(cors.New(config))

	userRoutes := router.Group("/users")
	routes.UserRoutes(userRoutes)

	authRequired := router.Group("/")
	authRequired.Use(middleware.Authentication())
	routes.UserRoutes(authRequired)

	//ADMIN USE
	router.POST("admin/playlists/createToDB", routes.AdminAddPlaylistToDB)
	router.POST("admin/food/createToDB", routes.AdminAddFoodToDB)
	router.POST("admin/restaurants/createToDB", routes.AdminAddRestaurantToDB)

	// Get All items
	router.GET("admin/playlists", routes.AdminGetPlaylists)
	router.GET("admin/food", routes.AdminGetFood)

	// USER USAGE
	// router.POST("/cart/createPlaylist", routes.createPlaylist)

	router.GET("/restaurants", routes.GetRestaurants)
	// GET /restaurantByCuisine
	router.GET("/restaurants/:restaurant_id", routes.GetFoodByRestaurantID) //need to test
	// router.GET("/playlist/:user_id", routes.GetPlaylistByUserID)
	router.GET("/cart/:user_id", routes.GetCartByUserID)
	// router.GET("/user", routes.GetUserByID)

	router.PUT("/restaurants/:food_id", routes.AddFoodItemToCart)
	//router.POST("/login", controller.Login)

	// router.DELETE("/cart/:food_id", routes.DeleteCartFoodItem)
	//get by id

	// U
	// router.PUT("/waiter/update/:id", routes.UpdateWaiter)

	// D
	// router.DELETE("/order/delete/:id", routes.DeleteOrder)

	//this runs the server and allows it to listen to requests.

	//FOR TESTDATA
	// testdata.DropTestData()
	// testdata.InsertData()

	router.Run(":" + port)
}
