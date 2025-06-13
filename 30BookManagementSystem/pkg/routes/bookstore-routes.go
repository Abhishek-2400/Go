package routes

import (
	"fmt"

	"github.com/Abhishek-2400/30BookManagementSystem/pkg/controllers"
	"github.com/gorilla/mux"
)

func init() {
	fmt.Println("temp")
}

func RegisterBookStoreRoutes(router *mux.Router) {
	fmt.Println("2")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	fmt.Println("23")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
