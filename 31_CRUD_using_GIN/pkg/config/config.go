package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() { //establish connection with db
	dsn := "host=localhost user=postgres password=arun196821803001 dbname=stocks port=5432 sslmode=disable"
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = dbInstance
}

func GetDB() *gorm.DB { //transfer the db string
	return db
}
