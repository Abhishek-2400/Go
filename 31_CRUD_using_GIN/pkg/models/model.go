package model

import (
	"github.com/Abhishek-2400/crud_gin/pkg/config"
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Owner string  `json:"owner"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Stock{})
}

func CreateStock(s *Stock) (*Stock, error) {
	result := db.Create(s)
	return s, result.Error
}

func GetAllStocks(allStocks *[]Stock) (*[]Stock, error) {
	result := db.Find(allStocks)
	return allStocks, result.Error
}

func GetStockById(id int64) (*Stock, error) {
	var stock Stock
	result := db.Where("ID=?", id).Find(&stock)
	// it's idiomatic in GORM and automatically returns ErrRecordNotFound if nothing is found.
	return &stock, result.Error
}

func DeleteStock(id int64) (*Stock, error) {
	var stock Stock
	result := db.Where("ID=?", id).Find(&stock)
	if result.Error != nil {
		return nil, result.Error
	}
	deleteError := db.Delete(&stock)
	return &stock, deleteError.Error
}
