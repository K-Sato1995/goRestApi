package main

import (
	"encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"log"
	// "math/rand"
	"net/http"
	// "strconv"
)

// Book Struct (Model)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Init Books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
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

	// Mock Data
	books = append(books, Book{ID: "1", Title: "Book one", Author: &Author{FirstName: "Philip", LastName: "Williams"}})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: &Author{FirstName: "John", LastName: "Johnson"}})

	// Route Hnadlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/", createBook).Methods("Book")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
