package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abhishek-2400/30BookManagementSystem/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println("0")
	routes.RegisterBookStoreRoutes(r) //routes is name of the modules Routes having func registerbookstore routes
	fmt.Println("1")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
