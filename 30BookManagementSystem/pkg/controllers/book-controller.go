package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Abhishek-2400/30BookManagementSystem/pkg/config"
	"github.com/Abhishek-2400/30BookManagementSystem/pkg/models"
	"github.com/Abhishek-2400/30BookManagementSystem/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func init() {
	fmt.Println("j")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var CreateBook = &models.Book{} // creates a pointer to an empty Book struct
	//var CreateBook = models.Book{}  // ceates an actual (value) Book struct, not a pointer
	utils.ParseBody(r, CreateBook) //json from web to our struct
	//utils.ParseBody(r, &CreateBook)
	b := models.CreateBook(CreateBook)
	//b := models.CreateBook(&CreateBook)

	res, _ := json.Marshal(b)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing")
	}
	bookDetails, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// ✅ Use the global db again, not the scoped one from GetBookById
	config.GetDB().Save(bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
