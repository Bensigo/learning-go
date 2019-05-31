package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/**
	book struct  model
  json field to be fetch
*/
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// init a book slice
var books []Book

// route helper function
func index(w http.ResponseWriter, r *http.Request) {

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r) // get params
	//loop through to find book with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
		json.NewEncoder(w).Encode(&Book{}) // return empyt ref of Book
	}

}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(book)
	book.ID = strconv.Itoa(rand.Intn(100000)) //convert to str
	books = append(books, book)
	// return the book
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["id"] {
			//	update the book
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	//return books
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["id"] {
			//delete
			books = append(books[:i], books[i+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(book)
			book.ID = params["id"]
			books = append(books, book)
			// return the book
			json.NewEncoder(w).Encode(book)

		}
	}
	//return books
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	//mock data
	books = append(books, Book{ID: "1", Isbn: "28383", Title: "Learn Python The Hard Way",
		Author: &Author{Firstname: "Zed", Lastname: "Shaw"}})
	books = append(books, Book{ID: "2", Isbn: "21233", Title: "ZERO TO ONE", Author: &Author{Firstname: "Peter", Lastname: "Thiel"}})
	// endpoint
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// run server
	log.Fatal(http.ListenAndServe(":8080", r))

}
