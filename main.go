package main

import (
	"rest/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", api.GetTodos)
	router.POST("/todos", api.AddTodo)
	router.Run("localhost:3000")
}