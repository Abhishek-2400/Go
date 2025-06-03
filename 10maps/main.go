package main

import "fmt"

func main() {
	fmt.Println("Maps in Go!")

	languages := make(map[string]string)

	languages["go"] = "Golang"
	languages["py"] = "Python"
	languages["js"] = "JavaScript"

	fmt.Printf("Languages map: %v\n", languages)
	fmt.Printf("Length of languages map: %d\n", len(languages))
	fmt.Printf("Value for key 'go': %s\n", languages["go"])

	delete(languages, "py")
	fmt.Printf("Languages map after deletion: %v\n", languages)
	

}