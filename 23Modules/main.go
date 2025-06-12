package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Modules")
	r := mux.NewRouter() //this is gorilla mux router
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page Hit")
	w.Write([]byte("Welcome to the Home Page!")) // Write response to the client in the form of bytes
}
