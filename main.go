package main

import (
	"read-books-library/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.RouterApp(router)
	router.Run(":8000")
}
