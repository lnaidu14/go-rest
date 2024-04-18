package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id        string `json:"id"`
	Title      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{
		Id: "1", Title: "Wake up", Completed: true,
	},
	{
		Id: "2", Title: "Brush teeth", Completed: false,
	},
	{
		Id: "3", Title: "Go for a jog", Completed: false,
	},
	{
		Id: "4", Title: "Eat breakfast", Completed: false,
	},
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func AddTodo(context *gin.Context) {
	var newTodo Todo
	if err := context.BindJSON(&newTodo) ; err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}