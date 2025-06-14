package routes

import (
	"github.com/Abhishek-2400/crud_gin/pkg/controllers"
	"github.com/Abhishek-2400/crud_gin/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func RequestRoute(r *gin.Engine) {
	r.GET("/stock", middlewares.CustomMW, controllers.GetAllStocks)
	r.GET("/stock/:stockid", controllers.GetStockById)
	r.POST("/stock", controllers.CreateStock)
	// r.PUT("/stock/{stockid}",controllers.UpdateStock)
	r.DELETE("/stock/:stockid", controllers.DeleteStock)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
}
