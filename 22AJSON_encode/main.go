package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	password string `json:"-"`                 // This field will be ignored in JSON output
	Address  string `json:"address,omitempty"` // This field will be omitted if empty
}

func main() {
	fmt.Println("Encode JSON") // struct -> json
	// JSON handling in Go is typically done using the "encoding/json" package.
	// This package provides functions to encode and decode JSON data.
	// You can use it to marshal Go structs into JSON format and unmarshal JSON data into Go structs.
	// Example of encoding a struct to JSON:

	persons := []Person{
		{"John", 30, "rrg", "123stret"},
		{"Jane", 25, "wef", "halbellstreetd31"},
		{"Doe", 40, "3ee", ""},
	}

	jsonDataByte, err := json.MarshalIndent(persons, "", " ")

	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON Data:%v", string(jsonDataByte))
}
