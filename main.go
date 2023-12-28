package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books/:id", bookHandler)
	v1.GET("/books", queryHandler)
	v1.POST("/books", postBooksHandler)

	router.Run(":8080")
}

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	SubTitle string `json:"subtitle" binding:"required"`
	Price    int    `json:"price" binding:"required"`
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		var errorMessages []string

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				errorMessage := fmt.Sprintf("error on field: %s, condition: %s", e.Field(), e.Tag())
				errorMessages = append(errorMessages, errorMessage)
			}
		} else {
			errorMessages = append(errorMessages, err.Error())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":    bookInput.Title,
		"subtitle": bookInput.SubTitle,
		"price":    bookInput.Price,
	})

}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}
