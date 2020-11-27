package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"math/rand"
	"strconv"
)
// Book Struct (Model)
type Book struct {
	ID		string	`json:"id"`
	Isbn	string	`json:"isbn"`
	Title	string	`json:"title"`
	Author	*Author	`json:"author"`
}

// * returns the value of the following variable .

//Author Struct (Model)
type Author struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(res).Encode(books)
}
// Get Single Book
func getBook(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Get params
	// loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			_ = json.NewEncoder(res).Encode(item)
			return // get out of here
		}
		//_ = json.NewEncoder(res).Encode(&Book{})
	}
}
// Create a New Book
func createbook(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	var book Book
	// & returns the memory address of the following variable.
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock Id - not safe
	books = append(books, book)
	_ = json.NewEncoder(res).Encode(book)
}

// Update book
func updateBook(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range books{
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) // slice
			res.Header().Set("Content-Type", "application/json")
			var book Book
			// & returns the memory address of the following variable.
			_ = json.NewDecoder(req.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			_ = json.NewEncoder(res).Encode(book)
			return
		}

	}
	_ = json.NewEncoder(res).Encode(books)
}
// Delete a new Book
func deleteBook(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range books{
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...) // slice
			break
		}

	}
	_ = json.NewEncoder(res).Encode(books)
}

func main(){
	// Init Router // type inference
	 router := mux.NewRouter()

	 // Mock data - @todo - implement DB
	books =  append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{
		FirstName: "Dickson",
		LastName:  "Douglas",
	} })
	books =  append(books, Book{ID: "2", Isbn: "478745", Title: "Book Two", Author: &Author{
		FirstName: "Smith",
		LastName:  "whales",
	} })

	 //  Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/book", createbook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook ).Methods("DELETE")

	 log.Fatal(http.ListenAndServe(":8000", router))
}