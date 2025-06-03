package main

import (
	"fmt"
	"io"
	"net/http"
)

// we will be using the net/http package to make web requests

const url = "https://google.com"

func main() {
	fmt.Println("Web requests in Go")

	response, err := http.Get(url) // this will send a GET request to the specified URL
	if err != nil {
		panic(err) // if there is an error, it will panic and stop the program
	}

	//fmt.Println("Response received :", response)
	defer response.Body.Close() // OUR RESPONSIBILTY TO CLOSE THE CONNECTION

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data received from the URL:", string(databytes))

}
