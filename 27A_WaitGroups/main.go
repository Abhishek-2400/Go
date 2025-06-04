package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signal = []string{"Hello"} //critical section
var mut sync.Mutex // Mutex to protect the critical section

func getStatusCode(website string) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes

	response, err := http.Get(website)
	if err != nil {
		fmt.Println("Error fetching website:", err)
	} else {
		fmt.Printf("Status code for %s: %d\n", website, response.StatusCode)
		mut.Lock() 
		signal = append(signal, website)
		mut.Unlock()
	}
}
func main() {
	fmt.Println("Wait Groups Example")
	// The problem with timer is that, our executors executes in bit amount of time so we unnecessarily  blocking the program for 1 sec.
	// Wait groups is a way of adding calculated and precise time delay  until all other goroutines have completed their execution.

	// 	WaitGroup exports 3 methods.
	// 1	Add(int)	 It increases WaitGroup counter by given integer value.
	// 2	Done()	     It decreases WaitGroup counter by 1, we will use it to indicate termination of a goroutine.
	// 3	Wait()	     It Blocks the execution until it's internal counter becomes 0.

	websites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
	}

	for _, website := range websites {
		go getStatusCode(website)
		wg.Add(1) // Increment the WaitGroup counter by 1 for each goroutine
	}

	wg.Wait() //// This Blocks the execution
	// until  WaitGroup counter become 0
	fmt.Println("signal:", signal)
}
