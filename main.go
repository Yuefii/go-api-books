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

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Book{})

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books/:id", handler.GetBookByID)
	v1.GET("/books", handler.GetAllBooks)
	v1.POST("/books", handler.PostBook)
	v1.PATCH("/books/:id", handler.UpdateBook)
	v1.DELETE("/books/:id", handler.DeleteBook)

	v1.GET("/categories", handler.GetAllCategory)
	v1.GET("/categories/:id", handler.GetCategoryByID)
	v1.POST("/categories", handler.PostCategory)
	v1.PATCH("/categories/:id", handler.UpdateCategory)
	v1.DELETE("/categories/:id", handler.DeleteCategory)

	auth := router.Group("/auth")

	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	router.Run(":8080")
}
