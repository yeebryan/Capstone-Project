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

	// Routes that don't require authentication
	router.GET("/restaurants", routes.GetRestaurants)
	router.GET("/restaurants/:restaurant_id", routes.GetFoodByRestaurantID) //need to test
	router.GET("/playlists", routes.GetPremadePlaylists)

	// USER USAGE
	// router.POST("/cart/createPlaylist", routes.createPlaylist)

	// Routes that require authentication

	{
		// router.GET("/playlist/:user_id", routes.GetPlaylistByUserID)
		authRequired.GET("/cart/:user_id", routes.GetCartByUserID)
		// router.GET("/user", routes.GetUserByID)
		authRequired.GET("/food/random", routes.FetchRandomFood)

		authRequired.PUT("/restaurants/:food_id", routes.AddFoodItemToCart)

		authRequired.GET("/playlists/:playlist_id", routes.GetFoodByPlaylistID)
		authRequired.GET("/playlists/food/:category", routes.GetFoodByCategory)
		authRequired.POST("/playlists/:playlist_id/create/:user_id", routes.CreateUserPremadePlaylist) //?start_date={start_date}
	}

	// router.DELETE("/cart/:food_id", routes.DeleteCartFoodItem)
	//get by id

	// U
	// router.PUT("/waiter/update/:id", routes.UpdateWaiter)

	// D
	// router.DELETE("/order/delete/:id", routes.DeleteOrder)

	//FOR TESTDATA
	// testdata.DropTestData()
	// testdata.InsertData()

	router.Run(":" + port)
}
