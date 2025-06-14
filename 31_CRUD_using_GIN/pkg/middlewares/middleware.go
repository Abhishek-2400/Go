package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CustomMW(c *gin.Context) {
	fmt.Println("custom Middleware check")
	c.Next()
}

// Function	Purpose
// c.Next()	Continues to the next middleware/handler
// c.Abort()	Stops the request chain immediately
