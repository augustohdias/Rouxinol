package main

import (
	"github.com/augustohdias/rouxinol/user"
	"github.com/augustohdias/rouxinol/user/repository"
	"log"
	"net/http"

	"github.com/augustohdias/rouxinol/db_client"
	"github.com/gorilla/mux"
)

//
//// GetPosts Retrieves all posts by user
//func GetPosts(w http.ResponseWriter, r *http.Request) {
//
//	userID := mux.Vars(r)["user_id"]
//	filter := bson.D{{Key: "user_id", Value: userID}}
//
//	collection := db.Collection("user_posts")
//	result, err := collection.Find(context.Background(), filter)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var posts []models.Post
//
//	result.All(context.Background(), &posts)
//	json.NewEncoder(w).Encode(posts)
//}
//
//// GetFeed Retrieves user feed
//func GetFeed(w http.ResponseWriter, r *http.Request) {
//	session := r.Header.Get("Session-Token")
//	userID := r.Header.Get("User-Id")
//
//	session = ""
//	userID = "0"
//	user := findUser(userID)
//
//	if session != user.SessionToken {
//		return
//	}
//
//	following := user.Following
//
//	filter := bson.D{
//		{Key: "user_id", Value: bson.D{
//			{Key: "$in", Value: following},
//		}},
//	}
//
//	collection := db.Collection("user_posts")
//	result, _ := collection.Find(context.Background(), filter)
//
//	var posts []models.Post
//
//	result.All(context.Background(), &posts)
//	json.NewEncoder(w).Encode(posts)
//}
//
//func findUser(userID string) *models.User {
//	collection := db.Collection("user")
//	result := collection.FindOne(context.Background(), bson.D{{Key: "user_id", Value: userID}})
//	var user models.User
//	result.Decode(&user)
//	return &user
//}

func main() {
	db := db_client.GetClient().Database("rouxinol")
	userRepository := repository.NewMongoDBRepository(db)
	postRepository := repository.
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
