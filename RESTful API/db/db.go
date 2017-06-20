package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=golang password=123123 sslmode=disable dbname=golang")

	if err != nil {
		panic(err)
	}

	return db
}	