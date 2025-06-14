package models

import (
	"fmt"

	"github.com/Abhishek-2400/30BookManagementSystem/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//Name string `gorm:"column:book_name"`
//“Hey GORM, even though this field is called Name in my Go struct,
// please map it to a column named book_name in the database.”

//gorm.Model
// This is an embedded struct provided by GORM. It adds the following standard fields to your Book model:
// type Model struct {
//     ID        uint           `gorm:"primaryKey"`
//     CreatedAt time.Time
//     UpdatedAt time.Time
//     DeletedAt gorm.DeletedAt `gorm:"index"`
// }

// AutoMigrate() → Ensures the table exists and has the correct columns
// &Book{} → Tells GORM: "Use this struct's definition to shape the table"

func CreateBook(b *Book) *Book {
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books) //find all the records based on given condition and populate the result in books struct
	//db.Where("author = ? AND publication = ?", "Alice", "O'Reilly").Find(&books)
	return Books
}

func GetBookById(Id int64) (*Book, error) {
	var getBook Book
	result := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, result.Error
}

func DeleteBook(ID int64) Book {
	var book Book
	result := db.Where("ID=?", ID).Delete(&book)
	if result.Error != nil {
		fmt.Println("Record not found or error:", result.Error)
		return Book{}
	}
	return book
}
