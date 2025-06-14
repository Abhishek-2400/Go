package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	model "github.com/Abhishek-2400/crud_gin/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllStocks(c *gin.Context) {
	var allStocks []model.Stock
	_, dbError := model.GetAllStocks(&allStocks)
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dbError.Error()})
	}
	c.JSON(http.StatusOK, allStocks)
}

func GetStockById(c *gin.Context) {
	fmt.Println("111")
	idStr := c.Param("stockid")
	fmt.Println(idStr)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}
	stockRequired, dbError := model.GetStockById(id)
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dbError.Error()})
	}
	c.JSON(http.StatusOK, stockRequired)
}

func CreateStock(c *gin.Context) {
	var newStock = &model.Stock{}
	error := c.ShouldBindJSON(newStock) //unmarshall req.body (json->struct) and then copy its content to newStock
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	stockCreated, dbError := model.CreateStock(newStock)
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dbError.Error()})
		return
	}
	c.JSON(http.StatusOK, stockCreated) //strct -> json
}

func DeleteStock(c *gin.Context) {
	idStr := c.Param("stockid")
	id, error := strconv.ParseInt(idStr, 10, 64)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}
	deletedStock, dbError := model.DeleteStock(id)
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dbError.Error()})
		return
	}
	c.JSON(http.StatusOK, deletedStock) //strct -> json

}

func Signup(c *gin.Context) {
	// get email/password

	var body = model.User{}
	error := c.BindJSON(&body)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	//hash password
	hash, errorPassword := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if errorPassword != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorPassword.Error(),
		})
		return
	}

	//create user
	body.Password = string(hash)

	user, dbError := model.Signup(&body)
	if dbError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": dbError.Error(),
		})
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	//get email/passowrd
	var user = model.User{}
	error := c.BindJSON(&user)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	fmt.Printf("user passwird %v\n", user.Password)
	token, jwtError := model.Login(&user)

	if jwtError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": jwtError.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
