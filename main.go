package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	db := Connect2DB()

	err := db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal(err)
	}

	SetupRoutes(engine, db)

	engine.Run(":8080")
}
