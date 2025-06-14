package model

import (
	"os"
	"time"

	"github.com/Abhishek-2400/crud_gin/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Stocks   []Stock `gorm:"foreignKey:UserID"` // 💡 One-to-many
}
type Stock struct {
	gorm.Model
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Owner  string  `json:"owner"`
	UserID uint    `json:"user_id"` // 🔗 Foreign key to User
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Stock{})
	db.AutoMigrate(&User{})
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

func Signup(u *User) (*User, error) {
	result := db.Create(u)
	return u, result.Error
}

func Login(u *User) (string, error) {
	var userFromDB User
	result := db.Where("email = ?", u.Email).First(&userFromDB)

	if result.Error != nil {
		return "", result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(u.Password))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": userFromDB.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
