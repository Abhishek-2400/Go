package main

import (
	"fmt"
	"time"
)

func greet(message string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(message)
	}
}

func main() {
	fmt.Println("Go Routines Example")
	go greet("Hello")
	greet("World")
}

// if we remove slpeep from the greet function, the program will exit immediately without printing "Hello" because the main goroutine will finish executing before the goroutine created by `go greet("Hello")` has a chance to run.
// Goroutines need scheduling opportunities to  run concurrently.

// when we put sleep scheduler gets enought time to context switch between the goroutines, allowing both "Hello" and "World" to be printed.

