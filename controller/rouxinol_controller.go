package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/augustohdias/Rouxinol/models"

	"github.com/augustohdias/Rouxinol/post"
	"github.com/augustohdias/Rouxinol/user"
	"github.com/gorilla/mux"
)

// RouxinolController encapsulates rouxinol's repositories
type RouxinolController struct {
	postRepo post.Repository
	userRepo user.Repository
}

// NewRouxinolController provides a controller
func NewRouxinolController(router *mux.Router, postRepo post.Repository, userRepo user.Repository) {
	handler := RouxinolController{postRepo, userRepo}
	router.HandleFunc("/user_posts/{username}", handler.GetPosts).Methods(http.MethodGet)
	router.HandleFunc("/user_posts", handler.NewPost).Methods(http.MethodPost)
}

// GetPosts retrieves all posts by user id
func (rx *RouxinolController) GetPosts(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["username"]
	posts, _ := rx.postRepo.GetAllByUsername(context.Background(), userID)
	json.NewEncoder(w).Encode(posts)
}

// NewPost creates a new post
func (rx *RouxinolController) NewPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var post models.Post
	err := decoder.Decode(&post)

	post.ID = post.Username + "_" + strconv.FormatInt(time.Now().UnixNano(), 16)

	fmt.Printf("New post for user %s Id: %s\n", post.Username, post.ID)

	if err != nil {
		msg := fmt.Sprintf("Request body contains badly-formed JSON")
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	rx.postRepo.New(context.Background(), post)
}
