package testdata

import (
	"context"
	"server/database"
	"server/models"
	"server/routes"
	"time"

	"github.com/xorcare/pointer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type foodData struct {
	ID    primitive.ObjectID
	Name  string
	Price float64
}

// Insert test data into the respective collections
func InsertData() error {
	// Collections
	foodCollection := database.OpenCollection(database.Client, "food")
	restaurantCollection := database.OpenCollection(database.Client, "restaurant")
	cartCollection := database.OpenCollection(database.Client, "cart")
	orderCollection := database.OpenCollection(database.Client, "order")

	// _id of respective data
	userID := make([]primitive.ObjectID, 3)
	restaurantID := make([]primitive.ObjectID, 10)

	// food data
	foodData := make([]foodData, 19)

	// User collection
	userData := []models.User{
		{
			FirstName:   pointer.String("Bryan"),
			LastName:    pointer.String("fp"),
			Email:       pointer.String("b@panda.com"),
			Password:    pointer.String("p12345"),
			Address:     "88 Marina Bay",
			PhoneNumber: pointer.String("8000-5555"),
		},
		{
			FirstName:   pointer.String("Daniel"),
			LastName:    pointer.String("fp"),
			Email:       pointer.String("d@panda.com"),
			Password:    pointer.String("p12345"),
			Address:     "88 Orchard",
			PhoneNumber: pointer.String("8001-5555"),
		},
		{
			FirstName:   pointer.String("Chek Wee"),
			LastName:    pointer.String("fp"),
			Email:       pointer.String("cw@panda.com"),
			Password:    pointer.String("p12345"),
			Address:     "88 Holland",
			PhoneNumber: pointer.String("8002-5555"),
		},
	}
	// Insert each user into the collection
	for i, user := range userData {
		createdUser := routes.AdminSignUp(user)
		userID[i] = createdUser.ID
	}

	// Food collection
	foodCollectionData := []models.Food{
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
	for i, food := range foodCollectionData {
		result, err := foodCollection.InsertOne(context.Background(), food)
		if err != nil {
			return err
		}
		foodData[i].ID = result.InsertedID.(primitive.ObjectID)
		foodData[i].Name = food.Name
		foodData[i].Price = food.Price
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
				foodData[0].ID,
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "Pizzeria del Mondo",
			Address:    "456 Elm St, Anytown USA",
			Categories: []string{"Pizza", "Italian"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodData[1].ID,
				foodData[2].ID,
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "Sushi Palace",
			Address:    "789 Oak St, Anytown USA",
			Categories: []string{"Sushi", "Japanese"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodData[3].ID,
				foodData[4].ID,
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "Taco Truck",
			Address:    "321 Maple St, Anytown USA",
			Categories: []string{"Mexican", "Tacos"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodData[5].ID,
				foodData[6].ID,
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "The Golden Wok",
			Address:    "555 Pine St, Anytown USA",
			Categories: []string{"Chinese", "Asian"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodData[7].ID,
				foodData[8].ID,
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "The Sizzling Skillet",
			Address:    "888 Oak St, Anytown USA",
			Categories: []string{"American", "Breakfast"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodData[9].ID,
				foodData[10].ID,
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "The Spice House",
			Address:    "222 Maple St, Anytown USA",
			Categories: []string{"Indian", "Spicy"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodData[11].ID,
				foodData[12].ID,
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "La Taqueria",
			Address:    "777 Main St, Anytown USA",
			Categories: []string{"Mexican", "Tacos"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodData[13].ID,
				foodData[14].ID,
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "Noodle House",
			Address:    "444 Elm St, Anytown USA",
			Categories: []string{"Asian", "Noodles"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodData[15].ID,
				foodData[16].ID,
			},
		},
		{
			ID:         primitive.NewObjectID(),
			Name:       "The Green Garden",
			Address:    "333 Oak St, Anytown USA",
			Categories: []string{"Vegetarian", "Healthy"},
			Image:      &models.ImageData{URL: "https://picsum.photos/seed/picsum/400/400"},
			Menu: []primitive.ObjectID{
				foodData[17].ID,
				foodData[18].ID,
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

	// Cart collection
	cartData := []models.Cart{
		{
			ID:     primitive.NewObjectID(),
			UserID: userID[0],
			Items: &[]models.FoodItems{
				{
					ID:       foodData[0].ID,
					Name:     foodData[0].Name,
					Quantity: 2,
					Price:    foodData[0].Price,
				},
				{
					ID:       foodData[1].ID,
					Name:     foodData[1].Name,
					Quantity: 1,
					Price:    foodData[1].Price,
				},
			},

			TotalPrice: 32.97,
			CreatedAt:  time.Date(2022, 2, 10, 8, 35, 0, 0, time.UTC),
		},
		{
			ID:     primitive.NewObjectID(),
			UserID: userID[0],
			Items: &[]models.FoodItems{
				{
					ID:       foodData[4].ID,
					Name:     foodData[4].Name,
					Quantity: 1,
					Price:    foodData[4].Price,
				},
				{
					ID:       foodData[2].ID,
					Name:     foodData[2].Name,
					Quantity: 2,
					Price:    foodData[2].Price,
				},
				{
					ID:       foodData[3].ID,
					Name:     foodData[3].Name,
					Quantity: 1,
					Price:    foodData[3].Price,
				},
			},
			TotalPrice: 49.96,
			CreatedAt:  time.Date(2022, 3, 5, 18, 20, 0, 0, time.UTC),
		},
		{
			ID:     primitive.NewObjectID(),
			UserID: userID[1],
			Items: &[]models.FoodItems{
				{
					ID:       foodData[5].ID,
					Name:     foodData[5].Name,
					Quantity: 1,
					Price:    foodData[5].Price,
				},
				{
					ID:       foodData[6].ID,
					Name:     foodData[6].Name,
					Quantity: 2,
					Price:    foodData[6].Price,
				},
			},
			TotalPrice: 10.97,
			CreatedAt:  time.Date(2022, 2, 28, 13, 45, 0, 0, time.UTC),
		},
	}
	// Insert each cart into the collection
	for _, cart := range cartData {
		_, err := cartCollection.InsertOne(context.Background(), cart)
		if err != nil {
			return err
		}
	}

	// Order collection
	orderData := []models.Order{
		{
			ID:     primitive.NewObjectID(),
			UserID: userID[0],
			Items: &[]models.FoodItems{
				{
					ID:       foodData[0].ID,
					Name:     foodData[0].Name,
					Quantity: 2,
					Price:    foodData[0].Price,
				},
				{
					ID:       foodData[1].ID,
					Name:     foodData[1].Name,
					Quantity: 1,
					Price:    foodData[1].Price,
				},
			},
			TotalPrice:    32.97,
			PaymentMethod: "credit",
			DeliveryTime:  time.Date(2022, 3, 22, 15, 0, 0, 0, time.UTC),
			CreatedAt:     time.Date(2022, 3, 12, 8, 0, 0, 0, time.UTC),
		},
		{
			ID:     primitive.NewObjectID(),
			UserID: userID[1],
			Items: &[]models.FoodItems{
				{
					ID:       foodData[2].ID,
					Name:     foodData[2].Name,
					Quantity: 1,
					Price:    foodData[2].Price,
				},
				{
					ID:       foodData[3].ID,
					Name:     foodData[3].Name,
					Quantity: 2,
					Price:    foodData[3].Price,
				},
			},
			TotalPrice:    32.97,
			PaymentMethod: "cash",
			DeliveryTime:  time.Date(2022, 3, 13, 20, 0, 0, 0, time.UTC),
			CreatedAt:     time.Date(2022, 3, 11, 18, 30, 0, 0, time.UTC),
		},
		{
			ID:     primitive.NewObjectID(),
			UserID: userID[2],
			Items: &[]models.FoodItems{
				{
					ID:       foodData[4].ID,
					Name:     foodData[4].Name,
					Quantity: 1,
					Price:    foodData[4].Price,
				},
				{
					ID:       foodData[5].ID,
					Name:     foodData[5].Name,
					Quantity: 1,
					Price:    foodData[5].Price,
				},
				{
					ID:       foodData[6].ID,
					Name:     foodData[6].Name,
					Quantity: 2,
					Price:    foodData[6].Price,
				},
			},
			TotalPrice:    21.96,
			PaymentMethod: "debit",
			DeliveryTime:  time.Date(2022, 3, 14, 18, 0, 0, 0, time.UTC),
			CreatedAt:     time.Date(2022, 3, 10, 10, 45, 0, 0, time.UTC),
		},
	}
	// Insert each order into the collection
	for _, order := range orderData {
		_, err := orderCollection.InsertOne(context.Background(), order)
		if err != nil {
			return err
		}
	}

	return nil
}

func DropTestData() {
	userCollection := database.OpenCollection(database.Client, "user")
	foodCollection := database.OpenCollection(database.Client, "food")
	restaurantCollection := database.OpenCollection(database.Client, "restaurant")
	cartCollection := database.OpenCollection(database.Client, "cart")
	orderCollection := database.OpenCollection(database.Client, "order")

	userCollection.Drop(context.Background())
	foodCollection.Drop(context.Background())
	restaurantCollection.Drop(context.Background())
	cartCollection.Drop(context.Background())
	orderCollection.Drop(context.Background())
}
