package main

import (
	"fmt"
)

// Shape is an interface that defines a method Area
type Shape interface {
	Area() float64
	permeter() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (c Circle) permeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (r Rectangle) permeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	fmt.Println("Interfaces in Go")
	// In Go, an interface is a type that specifies  methods that a type must implement to belong to that interface.
	// Interfaces are Go's way of achieving abstraction and generalization.
	// If a type implements all the methods in the interface, it implicitly satisfies the interface.
	// You can’t create an instance of an interface directly, but you can make a variable of the interface type to store any value that has the needed methods.

	//declare the variable 'shape' as type 'Shape' (the interface), not as 'Circle', so it can hold any value that implements the Shape interface.
	var shape Shape

	shape = Circle{Radius: 5}
	fmt.Printf("Circle Area: %v\n", shape.Area())
	fmt.Printf("Circle Perimeter: %v\n", shape.permeter())

	shape = Rectangle{Width: 4, Height: 6}
	fmt.Printf("Rectangle Area: %v\n", shape.Area())
	fmt.Printf("Rectangle Perimeter: %v\n", shape.permeter())

}

//Interface can hold any value, hence its   actual value and  type are determined dynamically in runtime
