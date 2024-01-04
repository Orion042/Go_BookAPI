package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(engine *gin.Engine, db *gorm.DB) {
	engine.POST("/book", func(c *gin.Context) {
		CreateBook(c, db)
	})

	engine.GET("/books", func(c *gin.Context) {
		GetAllBooks(c, db)
	})

	engine.GET("/book/:id", func(c *gin.Context) {
		GetBook(c, db)
	})

	engine.PUT("/book/:id", func(c *gin.Context) {
		UpdateBook(c, db)
	})

	engine.DELETE("/book/:id", func(c *gin.Context) {
		DeleteBook(c, db)
	})
}

func CreateBook(c *gin.Context, db *gorm.DB) {
	book := Book{}
	c.BindJSON(&book)

	if err := db.Create(&book).Error; err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, book)
}

func GetAllBooks(c *gin.Context, db *gorm.DB) {
	var books []Book

	if err := db.Find(&books).Error; err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context, db *gorm.DB) {
	book := Book{}
	id := c.Param("id")

	if err := db.First(&book, id).Error; err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context, db *gorm.DB) {
	book := Book{}
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		fmt.Println(err)
	}

	c.BindJSON(&book)
	db.Save(&book)

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context, db *gorm.DB) {
	book := Book{}
	id := c.Param("id")

	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		fmt.Println(err)
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, book)
}
