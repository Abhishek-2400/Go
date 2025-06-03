package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://localhost:8080/about?name=John&age=30"

func main() {
	fmt.Println("URLS in Go")
	result, _ := url.Parse(myURL)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("Raw Query:", result.RawQuery)

	qparams := result.Query()

	for key, value := range qparams {
		fmt.Printf("Query Parameter: %s = %s\n", key, value)
	}

	//constructing a new URL
	newURL := &url.URL{  //always pass a ref in this case
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/path/to/resource",
		RawQuery: "param1=value1&param2=value2",
	}

	fmt.Println("Constructed URL:", newURL.String())

}
