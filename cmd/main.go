package main

import (
	"github.com/dawood-usman/gorest/config"
	"github.com/dawood-usman/gorest/handlers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.MigrateDB()
}

func main() {

	r := gin.Default()

	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBookByID)
	r.POST("/books", handlers.CreateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)
	r.PUT("/books/:id", handlers.UpdateBook)

	r.Run(":8080")
}
