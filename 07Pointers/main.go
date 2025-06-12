package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go!")

	var num = 10

	//var ptr *int = &num         way 1
	//ptr:=num                    way 2
	var ptr = &num //             way 3  ptr is a pointer to num, it holds the memory address of num

	fmt.Printf("type is %T\n", ptr)
	fmt.Println("Value of num:", *ptr)
	fmt.Println("Address of num:", ptr)

	//pointer ensures that we are working with the same variable and not on a copy of it
	*ptr = 20
	fmt.Println("Updated value of num:", *ptr) // This will print 20, as we changed the value using the pointer

}
