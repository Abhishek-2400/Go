package routes

import (
	"github.com/Abhishek-2400/cron/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RouteRequest(r *gin.Engine) {
	r.GET("/", controllers.ServeHome)
}
