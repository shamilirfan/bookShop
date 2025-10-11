package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Struct define
type Book struct {
	ID       int     `json:"id"` // It is called tag
	BookName string  `json:"bookName"`
	Author   string  `json:"author"`
	Price    float32 `json:"price"`
	IsStock  bool    `json:"isStock"`
	ImageUrl string  `json:"imageUrl"`
}

// Book list define
var bookList []Book

// Port
var PORT string = ":4000"

// Create handlers
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Cors handling
	w.Header().Set("Access-Control-Allow-Origin", "*") // * means any IP address can access
	w.Header().Set("Content-Type", "application/json")

	// Write header
	w.WriteHeader(http.StatusOK)

	// Encoding
	json.NewEncoder(w).Encode(&bookList)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	// Cors handling
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Path Value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		http.Error(w, "Please give me a valid id", http.StatusBadRequest)
		return
	}

	// Store specific book
	var book Book

	// searching book
	for _, value := range bookList {
		if id == value.ID {
			book = value

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&book)
			return
		}
	}

	// Book not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Book Not Found!!!")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	// Cors handling
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Store new book
	var newBook Book

	// Decode
	decoder := json.NewDecoder(r.Body).Decode(&newBook)

	// Error handling
	if decoder != nil {
		http.Error(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	// Write header
	w.WriteHeader(http.StatusCreated)

	// Write a new book's ID
	newBook.ID = len(bookList) + 1

	// Append new book in a book list
	bookList = append(bookList, newBook)

	// Encode
	json.NewEncoder(w).Encode(&newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	// Cors handling
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Path value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		http.Error(w, "Please give me a valid id", http.StatusBadRequest)
		return
	}

	// Store updated book
	var updatedBook Book

	// Decode
	decoder := json.NewDecoder(r.Body).Decode(&updatedBook)

	// Error handling
	if decoder != nil {
		http.Error(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	// Searching specific book
	for index, value := range bookList {
		if id == value.ID {
			updatedBook.ID = id
			bookList[index] = updatedBook
		}
	}

	// Write header
	w.WriteHeader(http.StatusOK)

	// Encode and show update message
	json.NewEncoder(w).Encode("Successfully updated")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Cors handling
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Path value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		http.Error(w, "Please give a valid id", http.StatusBadRequest)
		return
	}

	// Store unmatched value
	var tempList []Book

	// Searching specific value
	for _, value := range bookList {
		if id != value.ID {
			tempList = append(tempList, value)
		}
	}

	bookList = tempList

	// Write header
	w.WriteHeader(http.StatusOK)

	// Encode and show delete message
	json.NewEncoder(w).Encode("Successfully deleted")
}

// Main function
func main() {
	// Create router
	mux := http.NewServeMux()

	// Create route/endpoint
	mux.Handle("GET /books", http.HandlerFunc(getBooks))
	mux.Handle("GET /books/{id}", http.HandlerFunc(getBook))
	mux.Handle("POST /books", http.HandlerFunc(createBook))
	mux.Handle("PUT /books/{id}", http.HandlerFunc(updateBook))
	mux.Handle("DELETE /books/{id}", http.HandlerFunc(deleteBook))

	// Listening server
	fmt.Println("Server is running at http://localhost" + PORT)
	serverListening := http.ListenAndServe(PORT, mux)

	// Handle Listening server error
	if serverListening != nil {
		fmt.Println("Something went wrong", serverListening)
	}
}

// Init function
func init() {
	// List of books
	books := []Book{
		{
			ID:       1,
			BookName: "The Promise of Heaven",
			Author:   "Dr. David Jeremiah",
			Price:    100,
			IsStock:  true,
			ImageUrl: "https://m.media-amazon.com/images/I/71agofkFeiL._SY466_.jpg",
		},
		{
			ID:       2,
			BookName: "How to Test Negative for Stupid",
			Author:   "John Kennedy",
			Price:    200,
			IsStock:  false,
			ImageUrl: "https://m.media-amazon.com/images/I/71tbImx2YVL._SY466_.jpg",
		},
		{
			ID:       3,
			BookName: "The Biography Of John Neely Kennedy",
			Author:   "Marcus Parker",
			Price:    300,
			IsStock:  false,
			ImageUrl: "https://m.media-amazon.com/images/I/61jC5-3L-XL._SY466_.jpg",
		},
		{
			ID:       4,
			BookName: "Last Rites",
			Author:   "Ozzy Osbourne",
			Price:    400,
			IsStock:  true,
			ImageUrl: "https://m.media-amazon.com/images/I/81L9X2TH++L._SY466_.jpg",
		},
		{
			ID:       5,
			BookName: "One Nation Always Under God",
			Author:   "Tim Scott",
			Price:    500,
			IsStock:  true,
			ImageUrl: "https://m.media-amazon.com/images/I/711h9Ts9CjL._SY466_.jpg",
		},
	}

	// Append book list in slice/list
	bookList = append(bookList, books...)
}
