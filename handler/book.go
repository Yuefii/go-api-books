package handler

import (
	"fmt"
	"net/http"

	"go-api-books/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func PostBooksHandler(c *gin.Context) {
	var bookInput models.BookInput

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
