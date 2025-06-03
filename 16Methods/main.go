package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

// You are defining a method on the user struct — so yes, it's like a member function of user, just like in object-oriented programming (OOP).
func (u User) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", u.Name, u.Age)
}
func (u *User) changeName(newName string) {
	u.Name = newName
}

func main() {
	fmt.Println("Methods in Go")
	user := User{Name: "Alice", Age: 30}
	fmt.Println(user.Greet())
	user.changeName("Bob")
	fmt.Println(user.Greet())
}
