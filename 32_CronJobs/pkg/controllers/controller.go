package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hi I am main being called by a cron job",
	})
}
