package routes

import (
	"context"
	"fmt"
	"net/http"
	"server/database"
	"server/models"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var playlistCollection *mongo.Collection = database.OpenCollection(database.Client, "playlist")

// add a playlist to db (admin side)
func AdminAddPlaylistToDB(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var playlist models.Playlist

	if err := c.BindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(playlist)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	playlist.ID = primitive.NewObjectID()

	result, insertErr := playlistCollection.InsertOne(ctx, playlist)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error adding playlist": insertErr.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// get all playlists (admin side)
func AdminGetPlaylists(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var playlists []bson.M

	cursor, err := playlistCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding playlist collection": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &playlists); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting playlist cursor": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, playlists)
}

// get all premade playlists
func GetPremadePlaylists(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var playlists []bson.M

	cursor, err := playlistCollection.Find(ctx, bson.M{"user_id": nil})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding playlist collection": err.Error()})
		return
	}

	if err = cursor.All(ctx, &playlists); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting playlist cursor": err.Error()})
		return
	}

	c.JSON(http.StatusOK, playlists)
}

// get all playlist items by the playlist's id
func GetFoodByPlaylistID(c *gin.Context) {
	playlistID := c.Params.ByName("playlist_id")
	docID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting playlist ID": err.Error()})
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var playlist models.Playlist

	err = playlistCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&playlist)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding food by playlist ID": err.Error()})
		return
	}

	var food = make([]models.Food, len(playlist.FoodID))
	for i, foodID := range playlist.FoodID {
		err = foodCollection.FindOne(ctx, bson.M{"_id": foodID}).Decode(&food[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error finding food": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, food)
}

// Create user playlist based on premade playlist
func CreateUserPremadePlaylist(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	playlistID := c.Param("playlist_id")
	playlistOID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting playlist ID": err.Error()})
		return
	}

	userID := c.Param("user_id")
	userOID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting user ID": err.Error()})
		return
	}

	startDate := c.Query("start_date")

	//get premade playlist
	//--exclude those crossed []excludedFoodItems?
	//change playlist by adding in new primitiveOID,userID, startdate
	//insert
	//think about diff between startdate and time.Now() to get array of food to order?
	var playlist models.Playlist

	err = playlistCollection.FindOne(ctx, bson.M{"_id": playlistOID}).Decode(&playlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding food by restaurant ID": err.Error()})
		return
	}

	//remove excluded items if excluded fields not nil
	//playlist.FoodID
	playlist.ID = primitive.NewObjectID()
	playlist.UserID = userOID
	playlist.StartDate = startDate

	_, insertErr := playlistCollection.InsertOne(ctx, playlist)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error adding playlist": insertErr.Error()})
		return
	}

	c.JSON(http.StatusOK, "Inserted successfully!")
}

// delete a playlist given the playlist id
func DeletePlaylist(c *gin.Context) {
	playlistID := c.Params.ByName("playlist_id")
	docID, _ := primitive.ObjectIDFromHex(playlistID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	deleteResult, err := playlistCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if deleteResult.DeletedCount > 0 {
		fmt.Printf("Deleted playlist with ID: %s\n", playlistID)
		c.JSON(http.StatusOK, "Deleted playlist successfully!")
	} else {
		fmt.Printf("No playlist found with ID: %s\n", playlistID)
		c.JSON(http.StatusNotFound, "No playlist found with the provided ID")
	}
}

// create user DIY playlist
type CreatePlaylistRequest struct {
	PlaylistName string `json:"playlistName"`
	Category     string `json:"category"`
	FoodType     string `json:"foodType"`
	Interval     string `json:"interval"`
	StartDate    string `json:"startDate"`
	Time         string `json:"time"`
	FoodId       string `json:"foodId"`
}

func CreateUserDIYPlaylist(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var request CreatePlaylistRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// get params
	userID := c.Value("uid")
	userOID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting user OID": err.Error()})
		return
	}

	foodOID, err := primitive.ObjectIDFromHex(request.FoodId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting food OID": err.Error()})
		return
	}

	intervalConv, err := models.IntervalType(request.Interval)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting interval": err.Error()})
		return
	}

	// playlist creation
	playlist := models.Playlist{
		ID:             primitive.NewObjectID(),
		Name:           request.PlaylistName,
		FoodID:         []primitive.ObjectID{foodOID},
		UserID:         userOID,
		Status:         models.StateOngoing,
		StartDate:      request.StartDate,
		DeliveryTiming: request.Time,
		TimingInterval: intervalConv,
	}

	_, insertErr := playlistCollection.InsertOne(ctx, playlist)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error adding playlist": insertErr.Error()})
		return
	}

	// order creation
	var food models.Food
	err = foodCollection.FindOne(ctx, bson.M{"_id": foodOID}).Decode(&food)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding food by food ID": err.Error()})
		return
	}

	order := models.Order{
		ID: primitive.NewObjectID(),
		Items: &[]models.FoodItems{
			{
				ID:       foodOID,
				Name:     food.Name,
				Quantity: 1,
				Price:    food.Price,
			},
		},
		UserID:       userOID,
		Status:       models.StatePending,
		StartDate:    request.StartDate,
		DeliveryTime: request.Time,
		CreatedAt:    time.Now(),
	}

	_, insertErr = orderCollection.InsertOne(ctx, order)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error adding order": insertErr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "Inserted DIY playlist successfully!",
		"playlist_id": playlist.ID.Hex(),
	})
}

// retrieve user PLAYLIST DA one
func GetPlaylistsCurrentUser(c *gin.Context) {
	fmt.Println("GetPlaylistsCurrentUser called") // Add this log statement
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userID := c.Value("uid")
	userOID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting user OID": err.Error()})
		return
	}
	var result []bson.M

	pipeline := []bson.M{
		{"$match": bson.M{
			"user_id": userOID},
		},
	}
	cursor, err := playlistCollection.Aggregate(ctx, pipeline)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error finding playlists by user ID": err.Error()})
		return
	}

	if err = cursor.All(ctx, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error getting playlist cursor": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

type PlaylistDB struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	FoodID []struct {
		Oid primitive.ObjectID `bson:"$oid"`
	} `bson:"food_id"`
	UserID struct {
		Oid primitive.ObjectID `bson:"$oid"`
	} `bson:"user_id"`
	Status    string `bson:"status"`
	StartDate struct {
		Date time.Time `bson:"$date"`
	} `bson:"start_date"`
	DeliveryTiming struct {
		Date struct {
			NumberLong int64 `bson:"$numberLong,string"`
		} `bson:"$date"`
	} `bson:"delivery_timing"`
	TimingInterval string `bson:"timing_interval"`
}

// retrieve user's playlist
func GetUserPlaylists(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	userID := c.Query("userId")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID parameter"})
		return
	}

	userOID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting user ID"})
		return
	}

	cursor, err := playlistCollection.Find(ctx, bson.M{"user_id": userOID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var playlistsDB []PlaylistDB
	if err = cursor.All(ctx, &playlistsDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("PlaylistsDB: ", playlistsDB) // Log playlistsDB

	playlists := make([]models.Playlist, len(playlistsDB))
	counter := 0
	for i, p := range playlistsDB {
		foodIDs := make([]primitive.ObjectID, len(p.FoodID))
		for j, f := range p.FoodID {
			foodIDs[j] = f.Oid
			counter++ // Increment the counter in each iteration

		}

		fmt.Println("Processed Playlists:", playlists) // Print processed playlists

		playlists[i] = models.Playlist{
			ID:             p.ID,
			Name:           p.Name,
			FoodID:         foodIDs,
			UserID:         p.UserID.Oid,
			Status:         models.State(p.Status),
			StartDate:      p.StartDate.Date.Format(time.RFC3339),
			DeliveryTiming: time.Unix(p.DeliveryTiming.Date.NumberLong/1000, 0).Format(time.RFC3339),
			TimingInterval: models.Interval(p.TimingInterval),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"playlists":    playlists,
		"userId":       userID,         // Debugging line
		"userOID":      userOID,        // Debugging line
		"playlistsLen": len(playlists), // Debugging line
		"counter":      counter,        // Return the counter value

	})
}

// // update the order
// func UpdateOrder(c *gin.Context) {

// 	orderID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(orderID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	var order models.Order

// 	if err := c.BindJSON(&order); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	validationErr := validate.Struct(order)
// 	if validationErr != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 		fmt.Println(validationErr)
// 		return
// 	}

// 	result, err := orderCollection.ReplaceOne(
// 		ctx,
// 		bson.M{"_id": docID},
// 		bson.M{
// 			"dish":   order.Dish,
// 			"price":  order.Price,
// 			"server": order.Server,
// 			"table":  order.Table,
// 		},
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	defer cancel()

// 	c.JSON(http.StatusOK, result.ModifiedCount)
// }

// // delete an order given the id
// func DeleteOrder(c *gin.Context) {

// 	orderID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(orderID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	result, err := orderCollection.DeleteOne(ctx, bson.M{"_id": docID})

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	defer cancel()

// 	c.JSON(http.StatusOK, result.DeletedCount)

// }
