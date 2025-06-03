package main

import "fmt"

func main() {
	
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("Three")

	fmt.Println("Hello")

	mydefer()

	fmt.Println("End of main function")
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// Output: hello , 4 3 2 1 0 , end of main function ,  three two one 