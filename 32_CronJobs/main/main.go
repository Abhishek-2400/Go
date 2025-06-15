package main

import (
	"github.com/Abhishek-2400/cron/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.RouteRequest(r)
	r.Run()
}
