package main

import (
	"fmt"
	"time"
)


func write(ch chan int) {
    for i := 0; i < 4; i++ {
        ch <- i
        fmt.Println("successfully wrote", i, "to ch")
    }
    close(ch)
}

func main() {
	fmt.Println("Channels in Go")

	//channels are way goroutines can communicate with each other
	//wait group is just the manager of all goroutines , routinees are not aware of each other presence in it
    
	//real world example (concurrent web scraping):---> 
	// Each goroutine is fetching a different URL. and sends the result to a shared channel.
	// The main goroutine waits for all goroutines to finish before printing the results.



    //Deadlock:-->

	// ch := make(chan string)   creates an unbuffered channel.
	// ch <- "geeks1"
	// fmt.Println(<-ch)      Since no one is receiving when you're trying to send, the program hangs and Go runtime detects a deadlock. (Our receiver is below the sender)

    //Deadlock:--->

	// ch := make(chan string)   creates an unbuffered channel.
	// fmt.Println(<-ch) 
	// ch <- "geeks1"   
	//Receiving (<-ch) blocks until a value is sent on the channel.
     

	//above can be solved by using buffer channels 
	// ch := make(chan string,1)  
	// ch <- "geeks1" 
	// fmt.Println(<-ch)  

	// By default channels are unbuffered, which states that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) which are ready to receive the sent value ans vice versa.
	// Buffered channels allows to accept a limited number of values without a corresponding receiver for those values.
	// Buffered channel are blocked only when the buffer is full.

	
    // creates capacity of 2
    ch := make(chan int, 2)
    go write(ch)
    time.Sleep(2 * time.Second)
    for v := range ch {
        fmt.Println("read value", v, "from ch")
        time.Sleep(2 * time.Second)

    }

	
}

//closing of channel is done by the sender, not the receiver.
// When a channel is closed, it can no longer be sent to, but it can still be received from.
// If you try to send to a closed channel, it will panic.
