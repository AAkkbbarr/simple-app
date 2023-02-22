package main

import (
	"net/http"
	"simple-app/internal/controllers"
)

var (
	postController = controllers.NewPostController()
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/posts", postController.GetAll)
	mux.HandleFunc("/post", postController.GetOne)
	mux.HandleFunc("/post/delete", postController.DeletePost)
	mux.HandleFunc("/post/create", postController.CreatePost)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic("Server is not started")
	}

}
