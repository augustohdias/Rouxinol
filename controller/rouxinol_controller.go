package controller

import (
	"context"
	"encoding/json"
	"github.com/augustohdias/rouxinol/post"
	"github.com/augustohdias/rouxinol/user"
	"github.com/gorilla/mux"
	"net/http"
)

type RouxinolController struct {
	postRepo post.Repository
	userRepo user.Repository
}

func NewRouxinolController(router *mux.Router, postRepo post.Repository, userRepo user.Repository) {
	handler := RouxinolController{postRepo, userRepo}
	router.HandleFunc("/user_posts/{id}", handler.GetPosts).Methods(http.MethodGet)
}

// GetPosts retrieves all posts by user id
func (rx *RouxinolController) GetPosts(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	posts, _ := rx.postRepo.GetAllByUserID(context.Background(), userID)
	json.NewEncoder(w).Encode(posts)
}
