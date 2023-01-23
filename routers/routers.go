package routers

import (
	"read-books-library/handlers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.BookById)
	router.GET("/ratingLower", handlers.OrderByRatingDecreasing)
	router.GET("/ratingHigher", handlers.OrderByRatingIncreasing)
	router.GET("/rate/:rating", handlers.RatingsOrdered)
	router.GET("/pages", handlers.SumOfPages)
	router.POST("/books", handlers.CreateBook)
	router.PATCH("/lowerRating", handlers.LowerRating)
	router.PATCH("/raiseRating", handlers.RaiseRating)
	router.Run(":8000")
}
