package main

import (
	"go-api-books/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books/:id", handler.BookHandler)
	v1.GET("/books", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8080")
}
