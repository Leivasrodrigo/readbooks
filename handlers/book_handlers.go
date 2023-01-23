package handlers

import (
	"errors"
	"net/http"
	"read-books-library/entities"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func LowerRating(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Rating <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	book.Rating -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func RaiseRating(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	book.Rating += 1
	c.IndentedJSON(http.StatusOK, book)
}

func GetBookById(id string) (*entities.Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func SumOfPages(c *gin.Context) {
	var sum int
	for _, b := range books {
		if b.Read {
			sum += b.Pages
		}
	}
	c.IndentedJSON(http.StatusOK, sum)
}

func CreateBook(c *gin.Context) {
	var newBook entities.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, books)
}

func OrderByRatingDecreasing(c *gin.Context) {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Rating > books[j].Rating
	})
	c.IndentedJSON(http.StatusCreated, books)
}

func OrderByRatingIncreasing(c *gin.Context) {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Rating < books[j].Rating
	})
	c.IndentedJSON(http.StatusCreated, books)
}

func OrderRatings(rating int) []entities.Book {
	var ratings []entities.Book
	for i, b := range books {
		if b.Rating == rating {
			ratings = append(ratings, books[i])
		}
	}
	return ratings
}

func RatingsOrdered(c *gin.Context) {
	r := c.Param("rating")
	rate, err := strconv.Atoi(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	ratings := OrderRatings(rate)
	c.IndentedJSON(http.StatusOK, ratings)
}

var books = []entities.Book{
	{ID: "001143679", Title: "Aconteceu Naquele Verão", Author: "Tessa Bailey", Rating: 4, Read: true, Pages: 448},
	{ID: "001151521", Title: "A prisioneira do tempo", Author: "Kate Morton", Rating: 5, Read: true, Pages: 448},
	{ID: "001120102", Title: "O Urso e o Rouxinol", Author: "Katherine Arden, Neil Gaiman", Rating: 4, Read: true, Pages: 320},
	{ID: "001160623", Title: "As infinitas possibilidades do nunca", Author: "Juliana Dantas", Rating: 1, Read: true, Pages: 355},
	{ID: "001156289", Title: "Altered Carbon", Author: "Richard Morgan", Rating: 0, Read: false, Pages: 544},
	{ID: "001159875", Title: "As aventuras de Tintim - O caso girassol", Author: "Dan Brown", Rating: 5, Read: true, Pages: 64},
	{ID: "001156841", Title: "Anjos e demônios (Robert Langdon)", Author: "Dan Brown", Rating: 4, Read: true, Pages: 480},
	{ID: "001168247", Title: "Na mira do vampiro", Author: "Lopes dos Santos", Rating: 3, Read: true, Pages: 120},
}
