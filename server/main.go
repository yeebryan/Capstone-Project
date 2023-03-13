package main

import (
	"os"

	"server/routes"

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

	router.Use(cors.Default())

	//ADMIN USE
	router.POST("admin/playlist/createToDB", routes.AdminAddPlaylistToDB)
	router.POST("admin/food/createToDB", routes.AdminAddFoodToDB)
	router.POST("admin/restaurant/createToDB", routes.AdminAddRestaurantToDB)

	// Get All items
	router.GET("admin/playlists", routes.AdminGetPlaylists)
	router.GET("admin/food", routes.AdminGetFood)

	// USER USAGE
	// router.POST("/cart/createPlaylist", routes.createPlaylist)

	router.GET("/restaurants", routes.GetRestaurants)
	// GET /restaurantByCuisine
	router.GET("/food/:restaurant_id", routes.GetFoodByRestaurantID) //need to test
	// router.GET("/playlist/:user_id", routes.GetPlaylistByUserID)
	// router.GET("/cart/:user_id", routes.GetCartByUserID)
	// router.GET("/user", routes.GetUserByID)

	// router.POST("/restaurants/:food_id", routes.AddFoodItemToCart)
	// router.POST("/login", routes.userLogin)

	// router.DELETE("/cart/:food_id", routes.DeleteCartFoodItem)
	//get by id

	// U
	// router.PUT("/waiter/update/:id", routes.UpdateWaiter)

	// D
	// router.DELETE("/order/delete/:id", routes.DeleteOrder)

	//this runs the server and allows it to listen to requests.

	//FOR TESTDATA
	// testdata.InsertData()
	// testdata.DropTestData()

	router.Run(":" + port)
}
