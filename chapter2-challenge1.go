package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    Book{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
    Book{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee"},
    Book{ID: 3, Title: "1984", Author: "George Orwell"},
}

func main() {
    router := gin.Default()

    // Get All Books
    router.GET("/books", func(c *gin.Context) {
        c.JSON(http.StatusOK, books)
    })

    // Get Book by ID
    router.GET("/books/:id", func(c *gin.Context) {
        id := c.Param("id")
        for _, book := range books {
            if fmt.Sprint(book.ID) == id {
                c.JSON(http.StatusOK, book)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
    })

    // Add Book
    router.POST("/books", func(c *gin.Context) {
        var newBook Book
        if err := c.BindJSON(&newBook); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        books = append(books, newBook)
        c.JSON(http.StatusCreated, newBook)
    })

    // Update Book
    router.PUT("/books/:id", func(c *gin.Context) {
        id := c.Param("id")
        var updatedBook Book
        if err := c.BindJSON(&updatedBook); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        for i, book := range books {
            if fmt.Sprint(book.ID) == id {
                books[i] = updatedBook
                c.JSON(http.StatusOK, updatedBook)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
    })

    // Delete Book
    router.DELETE("/books/:id", func(c *gin.Context) {
        id := c.Param("id")
        for i, book := range books {
            if fmt.Sprint(book.ID) == id {
                books = append(books[:i], books[i+1:]...)
                c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
    })

    router.Run(":8080")
}
