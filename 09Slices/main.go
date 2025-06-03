package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Slices in Go!\n")

	fruits := make([]int, 4)
	fruits[1] = 10
	fruits = append(fruits, 1, 2) // Appending elements to the slice
	fmt.Printf("Fruits slice: %v\n", fruits)

	sort.Ints(fruits)
	fmt.Printf("Sorted Fruits slice: %v\n", fruits)
	fmt.Printf("Is sorted : %v\n", sort.IntsAreSorted(fruits))

	//deleting an element from the slice
	indexToDelete := 1
	fruits = append(fruits[:indexToDelete], fruits[indexToDelete+1:]...)
	fmt.Printf("Fruits slice after deletion: %v\n", fruits)

}
