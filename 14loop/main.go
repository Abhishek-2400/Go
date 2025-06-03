package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Go!")

	var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Printf("Day %d: %s\n", i+1, days[i])
	// }

	// for i:=range days{
	// 	fmt.Printf("Day %d: %s\n", i+1, days[i])
	// }

	for index, day := range days {
		if index == 6 {
			goto lco
		}
		fmt.Printf("Day %d: %s\n", index+1, day)
	}

	//label lco only marks the line immediately after it — it does not create a block that contains the next few lines.
lco:
	fmt.Println("Looping completed!")

	fmt.Println("Using break and continue in loops:")

}
