package main

import "fmt"

// ✅ Go's Rule for Visibility (Exported vs Unexported):

// Exported (public)	    Start name with uppercase
// Unexported (private)   	Start name with lowercase	



type Person struct {
	Name    string    //exported
	Age     int       //exported
	Address string    //exported
    phone string //unexported, not accessible outside this package
}

func main() {
	fmt.Println("Structs in Go!")
	john := Person{"John Doe", 30, "123 Main St", "123-456-7890"}
	fmt.Printf("Person struct: %+v\n", john)
	fmt.Println("Name:", john.Name)
}
