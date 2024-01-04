package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func Connect2DB() *gorm.DB {
	dsn := "user:pass@tcp(db:3306)/book_db?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
