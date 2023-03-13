package testdata

import (
	"context"
	"server/models"
	"server/routes"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert test data into the respective collections
func InsertData() error {
	// Collections
	userCollection := routes.OpenCollection(routes.Client, "user")
	foodCollection := routes.OpenCollection(routes.Client, "food")
	restaurantCollection := routes.OpenCollection(routes.Client, "restaurant")

	//_id of respective data
	userID := make([]primitive.ObjectID, 3)
	foodID := make([]primitive.ObjectID, 19)
	restaurantID := make([]primitive.ObjectID, 10)

	// User collection
	userData := []models.User{
		{
			ID:          primitive.NewObjectID(),
			Name:        "Bryan",
			Email:       "b@panda.com",
			Password:    "p123",
			Address:     "88 Marina Bay",
			PhoneNumber: "8000-5555",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Daniel",
			Email:       "d@panda.com",
			Password:    "p1234",
			Address:     "88 Orchard",
			PhoneNumber: "8001-5555",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Chek Wee",
			Email:       "cw@panda.com",
			Password:    "p12345",
			Address:     "88 Holland",
			PhoneNumber: "8002-5555",
		},
	}
	// Insert each user into the collection
	for i, user := range userData {
		result, err := userCollection.InsertOne(context.Background(), user)
		if err != nil {
			return err
		}
		userID[i] = result.InsertedID.(primitive.ObjectID)
	}

	// Food collection
	foodData := []models.Food{
		{
			ID:          primitive.NewObjectID(),
			Name:        "Hamburger",
			Description: "Our classic hamburger with all the fixings",
			Price:       9.99,
			Image:       &models.ImageData{URL: "https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"beef", "classic"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Margherita Pizza",
			Description: "Our classic pizza with fresh tomato sauce, mozzarella cheese, and basil",
			Price:       12.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"vegetarian", "classic"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Pepperoni Pizza",
			Description: "Our classic pizza with tomato sauce, mozzarella cheese, and pepperoni",
			Price:       14.99,
			Image:       &models.ImageData{URL: " https://example.com/images/pepperoni-pizza.jpg"},
			Tag:         []string{"meat", "classic"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "California Roll",
			Description: "Crab, avocado, and cucumber rolled in rice and seaweed",
			Price:       8.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"seafood", "vegetarian"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Spicy Tuna Roll",
			Description: "Tuna and spicy mayo rolled in rice and seaweed",
			Price:       10.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"seafood", "spicy"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Al Pastor Taco",
			Description: "Pork with pineapple and cilantro",
			Price:       2.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"meat", "spicy"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Vegetarian Taco",
			Description: "Grilled vegetables with avocado and salsa",
			Price:       3.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"vegetarian"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Kung Pao Chicken",
			Description: "Stir-fried chicken with peanuts, vegetables, and spicy sauce",
			Price:       12.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"meat", "spicy"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Vegetable Stir-Fry",
			Description: "Stir-fried mixed vegetables in a light sauce",
			Price:       9.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"vegetarian"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Pancakes",
			Description: "Three fluffy pancakes with butter and syrup",
			Price:       6.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"classic", "breakfast"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Breakfast Burrito",
			Description: "Scrambled eggs, cheese, and bacon wrapped in a tortilla",
			Price:       8.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"meat", "breakfast"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Chicken Tikka Masala",
			Description: "Grilled chicken in a creamy tomato sauce with spices",
			Price:       13.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"meat", "spicy"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Vegetable Curry",
			Description: "Mixed vegetables in a spicy curry sauce",
			Price:       11.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"vegetarian", "spicy"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Carne Asada Taco",
			Description: "Grilled steak with cilantro and onion",
			Price:       2.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"meat"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Vegetarian Burrito",
			Description: "Rice, beans, cheese, and vegetables wrapped in a tortilla",
			Price:       8.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"vegetarian"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Pad Thai",
			Description: "Stir-fried rice noodles with egg, peanuts, and choice of meat",
			Price:       11.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"meat", "spicy"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Vegetable Ramen",
			Description: "Ramen noodles in a vegetable broth with mixed vegetables",
			Price:       9.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"vegetarian"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Quinoa Bowl",
			Description: "Mixed greens, quinoa, avocado, and roasted vegetables with a citrus dressing",
			Price:       10.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"vegetarian", "healthy"},
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Sweet Potato Wrap",
			Description: "Roasted sweet potatoes, mixed greens, and hummus wrapped in a whole wheat tortilla",
			Price:       8.99,
			Image:       &models.ImageData{URL: " https://picsum.photos/seed/picsum/200/200"},
			Tag:         []string{"vegetarian", "healthy"},
		},
	}
	// Insert each food item into the collection
	for i, food := range foodData {
		result, err := foodCollection.InsertOne(context.Background(), food)
		if err != nil {
			return err
		}
		foodID[i] = result.InsertedID.(primitive.ObjectID)
	}

	// Restaurant collection
	restaurantData := []models.Restaurant{
		{
			ID:         primitive.NewObjectID(),
			Name:       "The Burger Joint",
			Address:    "123 Main St, Anytown USA",
			Categories: []string{"Burgers", "Fast Food"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[0],
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "Pizzeria del Mondo",
			Address:    "456 Elm St, Anytown USA",
			Categories: []string{"Pizza", "Italian"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[1],
				foodID[2],
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "Sushi Palace",
			Address:    "789 Oak St, Anytown USA",
			Categories: []string{"Sushi", "Japanese"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[3],
				foodID[4],
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "Taco Truck",
			Address:    "321 Maple St, Anytown USA",
			Categories: []string{"Mexican", "Tacos"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[5],
				foodID[6],
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "The Golden Wok",
			Address:    "555 Pine St, Anytown USA",
			Categories: []string{"Chinese", "Asian"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[7],
				foodID[8],
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "The Sizzling Skillet",
			Address:    "888 Oak St, Anytown USA",
			Categories: []string{"American", "Breakfast"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[9],
				foodID[10],
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "The Spice House",
			Address:    "222 Maple St, Anytown USA",
			Categories: []string{"Indian", "Spicy"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[11],
				foodID[12],
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "La Taqueria",
			Address:    "777 Main St, Anytown USA",
			Categories: []string{"Mexican", "Tacos"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[13],
				foodID[14],
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "Noodle House",
			Address:    "444 Elm St, Anytown USA",
			Categories: []string{"Asian", "Noodles"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[15],
				foodID[16],
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "The Green Garden",
			Address:    "333 Oak St, Anytown USA",
			Categories: []string{"Vegetarian", "Healthy"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodID[17],
				foodID[18],
			},
		},
	}
	// Insert each restaurant into the collection
	for i, restaurant := range restaurantData {
		result, err := restaurantCollection.InsertOne(context.Background(), restaurant)
		if err != nil {
			return err
		}
		restaurantID[i] = result.InsertedID.(primitive.ObjectID)
	}
	return nil
}

func DropTestData(){
	userCollection := routes.OpenCollection(routes.Client, "user")
	foodCollection := routes.OpenCollection(routes.Client, "food")
	restaurantCollection := routes.OpenCollection(routes.Client, "restaurant")

	userCollection.Drop(context.Background())
	foodCollection.Drop(context.Background())
	restaurantCollection.Drop(context.Background())
}