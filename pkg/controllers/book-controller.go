package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PranavMasekar/go-bookstore/pkg/models"
	"github.com/PranavMasekar/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Get all books from model.go file
	newBooks := models.GetAllBooks()
	// Convert them to json
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	// Status code 200
	w.WriteHeader(http.StatusOK)
	// Write the response in w
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// Getting body of request
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	// Get book by id from DB
	bookDetails, _ := models.GetBookById(ID)
	// Convert to json and write to response
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Creating struct variable
	CreateBook := &models.Book{}
	// Setting created book variable with the response body
	// parsing json to object
	utils.ParseBody(r, CreateBook)
	// Adding entry to DB
	b := CreateBook.CreateBook()
	// Writing response in json
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Retriving ID from request body
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// Deletes the entry from DB with given Id and return the object of deleted object
	book := models.DeleteBook(ID)
	// Convet to json
	res, _ := json.Marshal(book)
	// Writing response
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// parse json to object
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	// Get parameters
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// Get the book by id
	bookDetails, db := models.GetBookById(ID)
	// set the values of bookDetails object equal to updateBook which is provided in request of api
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
