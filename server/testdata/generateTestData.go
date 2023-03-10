package testdata

import (
	"context"
	"server/models"
	"server/routes"
)

// Create a function to insert restaurant data into the "restaurant" collection
func insertRestaurantData() error {
	// Get a handle to the "restaurant" collection
	collection := routes.OpenCollection(routes.Client, "restaurant")

	// Create a slice of sample restaurant data
	restaurantData := []models.Restaurant{
		{
			CuisineType:    "Italian",
			RestaurantName: "Pizzeria Uno",
			// Rating:         4.5,
			// Location:       "Chicago, IL",
			// PostalCode:     60610,
			// OpeningHours:   []string{"Mon-Fri 11am-10pm", "Sat-Sun 11am-11pm"},
			// Image: &models.ImageData{
			// 	URL: "https://example.com/pizzeria-uno.jpg",
			// },
		},
		{
			CuisineType:    "Mexican",
			RestaurantName: "El Famous Burrito",
			// Rating:         4.0,
			// Location:       "Chicago, IL",
			// PostalCode:     60618,
			// OpeningHours:   []string{"Mon-Sun 10am-10pm"},
			// Image: &models.ImageData{
			// 	URL: "https://example.com/el-famous-burrito.jpg",
			// },
		},
		// add more sample data as needed
	}
	// Insert each food item into the collection
	for _, restaurant := range restaurantData {
		_, err := collection.InsertOne(context.Background(), restaurant)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create a function to insert food data into the "foods" collection
func insertFoodData() error {
	// Get a handle to the "foods" collection
	collection := routes.OpenCollection(routes.Client, "food")

	result1 := models.Restaurant{}
	result2 := models.Restaurant{}
	result := routes.FindRestaurantByName("Pizzeria Uno")
	result.Decode(result1)
	result = routes.FindRestaurantByName("El Famous Burrito")
	result.Decode(result2)

	// Create a slice of sample food data
	foodData := []models.Food{
		{
			Name: "Pepperoni Pizza",
			// Tag:          "pizza",
			// Halal:        false,
			RestaurantID: result1.ID,
			// Image: &models.ImageData{
			// 	URL: "https://example.com/pepperoni-pizza.jpg",
			// },
		},
		{
			Name: "Beef Burrito",
			// Tag:          "burrito",
			// Halal:        false,
			RestaurantID: result2.ID,
			// Image: &models.ImageData{
			// 	URL: "https://example.com/beef-burrito.jpg",
			// },
		},
		// add more sample data as needed
	}
	// Insert each food item into the collection
	for _, food := range foodData {
		_, err := collection.InsertOne(context.Background(), food)
		if err != nil {
			return err
		}
	}

	return nil
}
