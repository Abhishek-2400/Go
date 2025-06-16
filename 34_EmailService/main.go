package main

import "gopkg.in/gomail.v2"

func main() {
	abc := gomail.NewMessage()
	abc.SetHeader("From", "t31396978@gmail.com")
	abc.SetHeader("To", "t124@gmail.com")
	abc.SetBody("Subject", "Test Email")
	abc.SetBody("text/plain", "Test bidy")

	a := gomail.NewDialer("smtp.gmail.com", 587, "t31396978@gmail.com", "Ashish@123")
	err := a.DialAndSend(abc)

	if err != nil {
		panic(err)
	}

}
