package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses,omitempty"` // This field will be omitted if empty
	// the fileds must be exported (capitlised)  to be accessible by the json package
}

func decodeJSON() {
	jsonData := []byte(`{"name":"John", "age":30,"courses":["Go", "Python"]}`)
	var person Person

	checkValid := json.Valid(jsonData)
	if checkValid {
		json.Unmarshal(jsonData, &person)
		fmt.Println(person)
		//fmt.Println("Name:", person.Name)
	} else {
		fmt.Println("Invalid JSON data")
	}

	//sometimes you want to add data to map instead of struct

	var personMap map[string]interface{} //here we are sure that keys are going to be string and values can be anything when we deal with webdata so we use interfaces
	json.Unmarshal(jsonData, &personMap)
	//fmt.Println("Person Map:", personMap)

	for key, value := range personMap {
		fmt.Printf("%s: %v\n", key, value)
	}

}

func main() {
	fmt.Println("Decode JSON") //json -> struct
	decodeJSON()
}
