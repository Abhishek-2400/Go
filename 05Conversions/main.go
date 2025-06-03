package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter rating for pizza\n")

	// this method of reading input is used when space-separated values are expected
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error converting input to float:", err)
	} else {
		fmt.Println("Updated rating with +1 is: ", numRating+1)
	}

	// this method of reading input is used when single values are expected
	var name string

	fmt.Println("Enter your name:")
	fmt.Scan(&name)

	fmt.Printf("Name: %s\n", name)
}
