package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Introduce Yourself\n")

	// this method of reading input is used when space-separated values are expected
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Input received:", input)

}
