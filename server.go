package main

import (
	"github.com/DutkinTE/go/controllers"
	"github.com/DutkinTE/go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	models.ConnectDatabase()

	server.GET("/books", controllers.FindBooks)
	server.POST("/books", controllers.CreateBook)
	server.PATCH("/books/:id", controllers.UpdateBook)
	server.DELETE("/books/:id", controllers.DeleteBook)

	server.Run(":8000")
}
