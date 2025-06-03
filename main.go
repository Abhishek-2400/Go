package main

import (
	"fmt"
	"strings"
)

// Print prints the given message to the console.
func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Type of conferenceName is %T\n", conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string // This is a slice of strings, which can grow dynamically

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		fmt.Println(&firstName) //this will print the memory address of firstName

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		remainingTickets -= userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation will be sent to %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		bookings = append(bookings, firstName+" "+lastName)

		var firstNames []string
		for _, booking := range bookings {
			var firstName = strings.Fields(booking)
			firstNames = append(firstNames, firstName[0])
		}

		fmt.Printf("The whole slice is %v\n", bookings)
		fmt.Printf("The length of the slice is %v\n", len(bookings))
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

	}

}
