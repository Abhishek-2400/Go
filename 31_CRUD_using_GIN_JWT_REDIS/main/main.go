package main

import (
	"log"
	"time"

	"github.com/Abhishek-2400/crud_gin/pkg/middlewares"
	"github.com/Abhishek-2400/crud_gin/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	r := gin.Default() //Initializes a Gin router using gin.Default()
	r.Use(middlewares.RateLimiter(1, time.Hour))
	routes.RequestRoute(r)
	r.Run() //Starts the server on default port :8080.

	//When you call router.Run() without any arguments, Gin automatically checks for the PORT environment variable internally and uses it if it's available.

	// So the flow :is
	// init() → godotenv.Load() → loads .env into system environment
	// router.Run() → internally calls os.Getenv("PORT") → finds the PORT value
	// Gin uses that port automatically

	//   if u want to supply th env values in your code ,
	//   first u need to load them in go's system env  using godotenv.Load()
	//   then by using os.Getenv("PORT") you can access them in code and use them but rember in case of
	//   gin router.run() handles everything as sated above

	//   var apiKey = os.Getenv("PORT")
	//   fmt.Println(apiKey)

}

// router.GET("/", serveHome)
// router.GET("/ping", handlePing)

// func handlePing(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong",
// 	})

// }

// func serveHome(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API!"})
// }
