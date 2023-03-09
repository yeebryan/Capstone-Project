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

	// C
	//create timing
	// Create one item
	router.POST("/playlist/create", routes.AddPlaylist)
	router.POST("/food/create", routes.AddFood)
	router.POST("/restaurant/create", routes.AddRestaurant)

	// R
	// Get All items
	router.GET("/playlists", routes.GetPlaylists)
	router.GET("/food", routes.GetFood)
	router.GET("/restaurants", routes.GetRestaurants)
	//get by id

	// U
	// router.PUT("/waiter/update/:id", routes.UpdateWaiter)

	// D
	// router.DELETE("/order/delete/:id", routes.DeleteOrder)

	//this runs the server and allows it to listen to requests.
	router.Run(":" + port)
}
