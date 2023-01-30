# readbooks
API where you can manage the books you've read and rate them.
  router.GET("/books", handlers.GetBooks) //shows the list of books in the list as a JSON
	router.GET("/books/:id", handlers.BookById) //You can search a book by it's ID
	router.GET("/ratingLower", handlers.OrderByRatingDecreasing) //Show the list of books ordered by rating (higher to lower)
	router.GET("/ratingHigher", handlers.OrderByRatingIncreasing) //Show the list of books ordered by rating (lower to higher)
	router.GET("/rate/:rating", handlers.RatingsOrdered) //shows all books with the rating passed in the URL
	router.GET("/pages", handlers.SumOfPages) //sum of read pages
	router.POST("/books", handlers.CreateBook) //creates a book. You must insert a JSON as in the body.json file
	router.PATCH("/lowerRating", handlers.LowerRating) //decrease the rating of a book by passing ?id=<id>
	router.PATCH("/raiseRating", handlers.RaiseRating) //increase the rating of a book by passing ?id=<id>
