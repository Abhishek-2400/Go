package main

import (
	"fmt"
)

func sumSlice(numbers []int) (int,string) {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum,"Hi Abhishek!"
}

func main() {
	nums := []int{10, 20, 30, 40, 50}
	result,message := sumSlice(nums)
	fmt.Printf("The sum of the slice is: %d and message is %v\n", result,message)
}
