package handler

import (
	"fmt"
	"net/http"

	"go-api-books/config"
	"go-api-books/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAllBooks(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	var books []models.Book
	query := c.Request.URL.Query()

	title := query.Get("title")
	price := query.Get("price")
	rating := query.Get("rating")

	db = db.Model(&models.Book{})
	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if price != "" {
		db = db.Where("price = ?", price)
	}
	if rating != "" {
		db = db.Where("rating = ?", rating)
	}

	db.Find(&books)

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	var book models.Book
	bookID := c.Param("id")

	if err := db.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
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
