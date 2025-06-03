package main

import (
	"fmt"
)

func main() {
	println("If-Else in Go!")

	// if-else statement
	x := 0
	if x > 0 {
		fmt.Println("x is positive")
	} else if x < 0 {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is zero")
	}

	if num := 3; num < 0 {
		fmt.Println("num is negative")
	} else {
		fmt.Println("num is non-negative")
	}

	// switch statement
	switch x {
	case 0:
		println("x is zero")
		fallthrough
	case 1, 2, 3:
		println("x is one, two, or three")
	default:
		println("x is something else")
	}

	//  If the current case is executed and contains a fallthrough statement, then the very next case will be executed automatically regardless of its condition.
	//  But this applies only to the immediate next case, not beyond that unless the next case also has a fallthrough.
}
