package main

import (
	"go-api-books/config"
	"go-api-books/handler"
	"go-api-books/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Book{})

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books/:id", handler.GetBookByID)
	v1.GET("/books", handler.GetAllBooks)
	v1.POST("/books", handler.PostBook)
	v1.PATCH("/books/:id", handler.UpdateBook)
	v1.DELETE("/books/:id", handler.DeleteBook)

	router.Run(":8080")
}
