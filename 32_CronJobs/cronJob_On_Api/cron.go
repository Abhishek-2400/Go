package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Abhishek-2400/cron/pkg/config"
)

func makeApiCall() {
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	c := config.StartCorn()
	c.AddFunc("@every 5s", makeApiCall)
	c.Start() // Ensure the cron starts if not done inside StartCorn()

	// Block the main goroutine so the cron keeps running
	select {}
	//we can use for{} infinite for loop also but it consumes more cpu than select{}
}
