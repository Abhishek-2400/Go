package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func performGetRequest() {
	response, err := http.Get("http://localhost:5500")
	checkError(err)

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status:", response.Status)
	fmt.Println("Content Length:", response.ContentLength)

	databyte, err := io.ReadAll(response.Body)
	checkError(err)

	// 2 ways to convert byte slice to string
	// 1. Using string() conversion
	// 2. Using strings.Builder for more efficient concatenation

	//using strings.Builder
	var responseString strings.Builder
	byteCount, _ := responseString.Write(databyte)
	fmt.Println("Number of Bytes Written:", byteCount) //same as Content Length
	fmt.Println("Response String:", responseString.String())

	// Using string() conversion
	fmt.Println("Response Body:", string(databyte))

}

func performPostRequest() {
	const myURL = "http://localhost:5500/post"

	//Go does not support inline JSON-like object literals as you might see in JavaScript.
	requestBody := strings.NewReader(`
		{
			"name": "John",
			"age": 30
		}
	`)

	response, err := http.Post(myURL, "application/json", requestBody)
	checkError(err)

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status:", response.Status)

	databyte, err := io.ReadAll(response.Body)
	checkError(err)

	fmt.Println("Response Body:", string(databyte))
}

//A URL-encoded payload is data sent from a form using the application/x-www-form-urlencoded content type.

// If a form sends:
// <input name="username" value="abhishek" />
// <input name="age" value="22" />

// It arrives like this in the request body like:
// username=abhishek&age=22

// That’s URL-encoded format — key-value pairs joined with &, and special characters (like spaces) encoded (e.g., %20 for space).

//express.urlencoded() helps convert this into:
// req.body = { username: 'abhishek', age: '22' }

func postFormRequest() {
	const myURL = "http://localhost:5500/postform"
	//fake form data
	urlValues := url.Values{
		"name": {"John"},
		"age":  {"30"},
		"city": {"New York"},
	}

	fmt.Println("Form Data:", urlValues.Encode()) // converts to URL-encoded format like: name=John&age=30&city=New+York
	response, err := http.PostForm(myURL, urlValues) // postform sends data as application/x-www-form-urlencoded automatically we dont need to encode it manually
	checkError(err)
	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)
	checkError(err)
	fmt.Println("Response body:", string(databyte))
	// response, err := http.PostForm(myURL,
}


func main() {
	fmt.Println("GET requests in Go")
	postFormRequest()
	performGetRequest()
	performPostRequest()
	
}
