package main

import (
	"github.com/augustohdias/Rouxinol/controller"
	"github.com/augustohdias/Rouxinol/mongo_client"
	_postRepo "github.com/augustohdias/Rouxinol/post/repository"
	_userRepo "github.com/augustohdias/Rouxinol/user/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := mongo_client.GetClient().Database("Rouxinol")
	userRepository := _userRepo.NewMongoDBUserRepository(db)
	postRepository := _postRepo.NewMongoDBPostRepository(db)

	router := mux.NewRouter()

	controller.NewRouxinolController(router, postRepository, userRepository)

	log.Fatal(http.ListenAndServe(":8000", router))
}
