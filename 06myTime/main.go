package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in Go!")

	// date ,time and number given in Fomat is fixed for go , may be foundation date of go
	var currTime = time.Now().Format("01-02-2006 15:04:05 Monday") //MM/DD/YYYY HH:MM:SS Weekday
	fmt.Println("Current time is:", currTime)

	var createdDate=time.Date(2023,time.April,12,6,25,30,100,time.Local).Format("01-02-2006 15:04:05 Monday") //MM/DD/YYYY HH:MM:SS Weekday
	fmt.Println("Created date is:", createdDate)

}

