package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"log"
	// "math/rand"
	"net/http"
	// "strconv"
)

// Book Struct (Model)

type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// Create a Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Initiate Router
	r := mux.NewRouter()

	// Route Hnadlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/", createBook).Methods("Book")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
