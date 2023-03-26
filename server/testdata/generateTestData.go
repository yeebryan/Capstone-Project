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
	playlistCollection := database.OpenCollection(database.Client, "playlist")

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
			Image:       &models.ImageData{URL: "https://i.ibb.co/xm7PvQZ/hamburger.jpg"},
			Category:    "Western",
			FoodType:    "Burger",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Margherita Pizza",
			Description: "Our classic pizza with fresh tomato sauce, mozzarella cheese, and basil",
			Price:       12.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/cvdjHmw/marg-pizza.jpg"},
			Category:    "Western",
			FoodType:    "Pizza",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Pepperoni Pizza",
			Description: "Our classic pizza with tomato sauce, mozzarella cheese, and pepperoni",
			Price:       14.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/3yWdpM1/pepporoni.jpg"},
			Category:    "Western",
			FoodType:    "Pizza",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "California Roll",
			Description: "Crab, avocado, and cucumber rolled in rice and seaweed",
			Price:       8.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/SNSjpw4/cali-roll.jpg"},
			Category:    "Japanese",
			FoodType:    "Sushi",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Salmon Roll",
			Description: "Salmon and spicy fish roe rolled in rice and seaweed",
			Price:       10.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/SRgswHF/sushi.jpg"},
			Category:    "Japanese",
			FoodType:    "Sushi",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Al Pastor Taco",
			Description: "Pork with pineapple and cilantro",
			Price:       2.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/Yb015Zp/taco-copy.jpg"},
			Category:    "Mexican",
			FoodType:    "Taco",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Plant-Based Taco",
			Description: "Grilled vegetables with plant-based chicken strips",
			Price:       3.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/zPwWXMS/taco-2.jpg"},
			Category:    "Mexican",
			FoodType:    "Taco",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Kung Pao Chicken",
			Description: "Stir-fried chicken with peanuts, vegetables, and spicy sauce",
			Price:       12.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/68GZx7G/kungfu-chicken.jpg"},
			Category:    "Chinese",
			FoodType:    "Chicken",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Chow Mein",
			Description: "Stir-fried Chow Mein in a light sauce",
			Price:       9.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/NyY8gP3/chow-mein.jpg"},
			Category:    "Chinese",
			FoodType:    "Noodle",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Pancakes",
			Description: "Three fluffy pancakes with butter and syrup",
			Price:       6.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/SJK3dy4/pancake.jpg"},
			Category:    "Western",
			FoodType:    "Breakfast",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Breakfast Burrito",
			Description: "Scrambled eggs, cheese, and bacon wrapped in a tortilla",
			Price:       8.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/k8Wzgyp/burrito-pork.jpg"},
			Category:    "Western",
			FoodType:    "Breakfast",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Chicken Tikka Masala",
			Description: "Grilled chicken in a creamy tomato sauce with spices",
			Price:       13.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/KVRQjgm/chicken-masala.jpg"},
			Category:    "Indian",
			FoodType:    "Chicken",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Vegetable Curry Rice",
			Description: "Mixed vegetables in a spicy curry sauce",
			Price:       11.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/0BWtP2m/indian-curry-rice.jpg"},
			Category:    "Indian",
			FoodType:    "Rice",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Carne Asada Taco",
			Description: "Grilled steak with cilantro and onion",
			Price:       2.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/r3CHXty/taco-3.jpg"},
			Category:    "Mexican",
			FoodType:    "Taco",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Vegetarian Burrito",
			Description: "Rice, beans, cheese, and vegetables wrapped in a tortilla",
			Price:       8.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/dQRtr1P/burrito.jpg"},
			Category:    "Mexican",
			FoodType:    "Rice",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Yangzhou Fried Rice",
			Description: "Stir-fried rice dish that is made with cooked rice, eggs, ham, shrimp, green onions, and other vegetables",
			Price:       11.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/7VPVtRV/fried-rice.jpg"},
			Category:    "Chinese",
			FoodType:    "Rice",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Vegetable Noodle",
			Description: "Yellow noodles in a vegetable broth with mixed vegetables",
			Price:       9.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/cgSJHth/yellow-noodle.jpg"},
			Category:    "Chinese",
			FoodType:    "Noodle",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Quinoa Bowl",
			Description: "Mixed greens, quinoa, avocado, and roasted vegetables with a citrus dressing",
			Price:       10.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/VxQYNFs/salad-2.jpg"},
			Category:    "Vegetarian",
			FoodType:    "Salad",
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Sweet Potato Wrap",
			Description: "Roasted sweet potatoes, mixed greens, and hummus wrapped in a whole wheat tortilla",
			Price:       8.99,
			Image:       &models.ImageData{URL: "https://i.ibb.co/QpJVMwX/wrap.jpg"},
			Category:    "Vegetarian",
			FoodType:    "Salad",
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
			ID:       primitive.NewObjectID(),
			Name:     "The Burger Joint",
			Address:  "123 Main St, Anytown USA",
			Category: "Western",
			Image:    &models.ImageData{URL: "https://i.ibb.co/K59sMZc/burger-joint.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[0].ID,
				},
			},
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "Pizzeria del Mondo",
			Address:  "456 Elm St, Anytown USA",
			Category: "Western",
			Image:    &models.ImageData{URL: "https://i.ibb.co/Bckz4tb/del-mondo.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[1].ID,
					foodData[2].ID,
				},
			},
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "Sushi Palace",
			Address:  "789 Oak St, Anytown USA",
			Category: "Japanese",
			Image:    &models.ImageData{URL: "https://i.ibb.co/SmVw8Qg/sushi.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[3].ID,
					foodData[4].ID,
				},
			},
		},

		{
			ID:       primitive.NewObjectID(),
			Name:     "Taco Truck",
			Address:  "321 Maple St, Anytown USA",
			Category: "Mexican",
			Image:    &models.ImageData{URL: "https://i.ibb.co/YjtLtfV/taco-truck.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[5].ID,
					foodData[6].ID,
				},
			},
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "The Golden Wok",
			Address:  "555 Pine St, Anytown USA",
			Category: "Chinese",
			Image:    &models.ImageData{URL: "https://i.ibb.co/qgkT4zx/golden-wok.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[7].ID,
					foodData[8].ID,
				},
			},
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "The Sizzling Skillet",
			Address:  "888 Oak St, Anytown USA",
			Category: "Western",
			Image:    &models.ImageData{URL: "https://i.ibb.co/Vqn1TzP/sizzling-skillet.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[9].ID,
					foodData[10].ID,
				},
			},
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "The Spice House",
			Address:  "222 Maple St, Anytown USA",
			Category: "Indian",
			Image:    &models.ImageData{URL: "https://i.ibb.co/fCpSvJ1/spice-house.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[11].ID,
					foodData[12].ID,
				},
			},
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "La Taqueria",
			Address:  "777 Main St, Anytown USA",
			Category: "Mexican",
			Image:    &models.ImageData{URL: "https://i.ibb.co/GtDwXPx/la-taqueria.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[13].ID,
					foodData[14].ID,
				},
			},
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "Noodle House",
			Address:  "444 Elm St, Anytown USA",
			Category: "Chinese",
			Image:    &models.ImageData{URL: "https://i.ibb.co/FX3DTtg/noodle-house.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[15].ID,
					foodData[16].ID,
				},
			},
		},
		{
			ID:       primitive.NewObjectID(),
			Name:     "The Green Garden",
			Address:  "333 Oak St, Anytown USA",
			Category: "Vegetarian",
			Image:    &models.ImageData{URL: "https://i.ibb.co/Rzzzk3v/green-garden.jpg"},
			Menu: models.Menu{
				[]primitive.ObjectID{
					foodData[17].ID,
					foodData[18].ID,
				},
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
			Items: []models.FoodItems{
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
			TotalPrice: foodData[0].Price + foodData[1].Price,
			CreatedAt:  time.Date(2022, 2, 10, 8, 35, 0, 0, time.UTC),
			State:      models.StateInProcess,
		},
		{
			ID:     primitive.NewObjectID(),
			UserID: userID[0],
			Items: []models.FoodItems{
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
				{
					ID:       foodData[4].ID,
					Name:     foodData[4].Name,
					Quantity: 1,
					Price:    foodData[4].Price,
				},
			},

			TotalPrice: (foodData[2].Price * 2) + foodData[3].Price + foodData[4].Price,
			CreatedAt:  time.Date(2022, 3, 5, 18, 20, 0, 0, time.UTC),
			State:      models.StateInProcess,
		},
		{
			ID:     primitive.NewObjectID(),
			UserID: userID[1],
			Items: []models.FoodItems{
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
			TotalPrice: foodData[5].Price + (foodData[6].Price * 2),
			CreatedAt:  time.Date(2022, 2, 28, 13, 45, 0, 0, time.UTC),
			State:      models.StateInProcess,
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
			TotalPrice:    (foodData[0].Price * 2) + foodData[1].Price,
			PaymentMethod: "credit",
			DeliveryTime:  time.Date(2022, 3, 22, 15, 0, 0, 0, time.UTC),
			CreatedAt:     time.Date(2022, 3, 12, 8, 0, 0, 0, time.UTC),
			Status:        models.StateCompleted,
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
			TotalPrice:    foodData[2].Price + (foodData[3].Price * 2),
			PaymentMethod: "cash",
			DeliveryTime:  time.Date(2022, 3, 13, 20, 0, 0, 0, time.UTC),
			CreatedAt:     time.Date(2022, 3, 11, 18, 30, 0, 0, time.UTC),
			Status:        models.StateCompleted,
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
			TotalPrice:    foodData[4].Price + foodData[5].Price + (foodData[6].Price * 2),
			PaymentMethod: "debit",
			DeliveryTime:  time.Date(2022, 3, 14, 18, 0, 0, 0, time.UTC),
			CreatedAt:     time.Date(2022, 3, 10, 10, 45, 0, 0, time.UTC),
			Status:        models.StateCompleted,
		},
	}
	// Insert each order into the collection
	for _, order := range orderData {
		_, err := orderCollection.InsertOne(context.Background(), order)
		if err != nil {
			return err
		}
	}

	// Playlist collection
	playlistData := []models.Playlist{
		{
			ID:   primitive.NewObjectID(),
			Name: "Merry Monday",
			FoodID: []primitive.ObjectID{
				foodData[0].ID,
				foodData[1].ID,
				foodData[2].ID,
				foodData[3].ID,
			},
			UserID: primitive.NilObjectID,
			Image:  &models.ImageData{URL: "https://i.ibb.co/N3s0QmR/3.png"},
		},
		{
			ID:   primitive.NewObjectID(),
			Name: "Totally Tuesday",
			FoodID: []primitive.ObjectID{
				foodData[4].ID,
				foodData[5].ID,
				foodData[6].ID,
				foodData[7].ID,
			},
			UserID: primitive.NilObjectID,
			Image:  &models.ImageData{URL: "https://i.ibb.co/1v590Fs/4.png"},
		},
		{
			ID:   primitive.NewObjectID(),
			Name: "Wonder Wednesday",
			FoodID: []primitive.ObjectID{
				foodData[8].ID,
				foodData[9].ID,
				foodData[10].ID,
				foodData[11].ID,
			},
			UserID: primitive.NilObjectID,
			Image:  &models.ImageData{URL: "https://i.ibb.co/MPLsPJy/5.png"},
		},
		{
			ID:   primitive.NewObjectID(),
			Name: "Tasty Thursday",
			FoodID: []primitive.ObjectID{
				foodData[12].ID,
				foodData[13].ID,
				foodData[14].ID,
				foodData[15].ID,
			},
			UserID: primitive.NilObjectID,
			Image:  &models.ImageData{URL: "https://i.ibb.co/vQ0XqYL/2.png"},
		},
		{
			ID:   primitive.NewObjectID(),
			Name: "Funky Friday",
			FoodID: []primitive.ObjectID{
				foodData[16].ID,
				foodData[17].ID,
				foodData[18].ID,
				foodData[0].ID,
			},
			UserID: primitive.NilObjectID,
			Image:  &models.ImageData{URL: "https://i.ibb.co/dBQRs3n/1.png"},
		},
		{
			ID:   primitive.NewObjectID(),
			Name: "Sweet Saturday",
			FoodID: []primitive.ObjectID{
				foodData[1].ID,
				foodData[5].ID,
				foodData[8].ID,
				foodData[11].ID,
			},
			UserID: primitive.NilObjectID,
			Image:  &models.ImageData{URL: "https://i.ibb.co/hynpxj5/6.png"},
		},
		{
			ID:   primitive.NewObjectID(),
			Name: "Superb Sunday",
			FoodID: []primitive.ObjectID{
				foodData[2].ID,
				foodData[10].ID,
				foodData[15].ID,
				foodData[17].ID,
			},
			UserID: primitive.NilObjectID,
			Image:  &models.ImageData{URL: "https://i.ibb.co/PF6Wf08/7.png"},
		},
	}
	// Insert each playlist into the collection
	for _, playlist := range playlistData {
		_, err := playlistCollection.InsertOne(context.Background(), playlist)
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
	playlistCollection := database.OpenCollection(database.Client, "playlist")

	userCollection.Drop(context.Background())
	foodCollection.Drop(context.Background())
	restaurantCollection.Drop(context.Background())
	cartCollection.Drop(context.Background())
	orderCollection.Drop(context.Background())
	playlistCollection.Drop(context.Background())
}
