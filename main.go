package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	// "math/rand"
	"net/http"
	// "strconv"
	"log"
)

// Post Struct (Model)

type Post struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Get All Posts
func getPosts(w http.ResponseWriter, r *http.Request) {

}

// Get Single Post
func getPost(w http.ResponseWriter, r *http.Request) {

}

// Create a Post
func createPost(w http.ResponseWriter, r *http.Request) {

}

// Update a Post
func updatePost(w http.ResponseWriter, r *http.Request) {

}

// Delete a Post
func deletePost(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Initiate Router
	r := mux.NewRouter()

	// Route Hnadlers / Endpoints
	r.HandleFunc("/api/posts", getPosts).Methods("GET")
	r.HandleFunc("/api/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/api/posts/", createPost).Methods("POST")
	r.HandleFunc("/api/posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/api/posts/{id}", deletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
