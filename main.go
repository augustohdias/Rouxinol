package main

import (
	"github.com/augustohdias/rouxinol/controller"
	"github.com/augustohdias/rouxinol/mongo_client"
	_postRepo "github.com/augustohdias/rouxinol/post/repository"
	_userRepo "github.com/augustohdias/rouxinol/user/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := mongo_client.GetClient().Database("rouxinol")
	userRepository := _userRepo.NewMongoDBUserRepository(db)
	postRepository := _postRepo.NewMongoDBPostRepository(db)

	router := mux.NewRouter()

	controller.NewRouxinolController(router, postRepository, userRepository)

	log.Fatal(http.ListenAndServe(":8000", router))
}
