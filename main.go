package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Book struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
}

func connect2db() *gorm.DB {
	dsn := "user:pass@tcp(db:3306)/book_db?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	engine := gin.Default()

	db := connect2db()

	err := db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal(err)
	}

	engine.POST("/book", func(c *gin.Context) {
		book := Book{}

		c.BindJSON(&book)

		if err := db.Create(&book).Error; err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, book)
	})

	engine.GET("/book/:id", func(c *gin.Context) {
		book := Book{}

		id := c.Param("id")

		if err := db.First(&book, id).Error; err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, book)
	})

	engine.PUT("/book/:id", func(c *gin.Context) {
		book := Book{}

		id := c.Param("id")

		if err := db.Where("id = ?", id).First(&book).Error; err != nil {
			fmt.Println(err)
		}

		c.BindJSON(&book)

		db.Save(&book)

		c.JSON(http.StatusOK, book)
	})

	engine.DELETE("/book/:id", func(c *gin.Context) {
		book := Book{}

		id := c.Param("id")

		if err := db.Where("id = ?", id).First(&book).Error; err != nil {
			fmt.Println(err)
		}

		db.Delete(&book)

		c.JSON(http.StatusOK, book)
	})

	engine.Run(":8080")
}
