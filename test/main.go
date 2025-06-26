package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Task: "Learn Go", Completed: false},
	{ID: "2", Task: "Build a REST API", Completed: false},
}

// get function
func getodo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)

}

// post
func posttodo(c *gin.Context) {
	var newTodo todo
	if er := c.BindJSON(&newTodo); er != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// serach
func search(id string) (*todo, error) {
	for _, t := range todos {
		if t.ID == id {
			return &t, nil

		}
	}
	return nil, errors.New("ttodo not found")
}

func serachtodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := search(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

// update
func updatetodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := search(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

func main() {

	Router := gin.Default()

	Router.GET("/todo", getodo)
	Router.POST("/todo", posttodo)
	Router.GET("/todo/:id", serachtodo)
	Router.PATCH("/todo/:id", updatetodo)
	Router.Run("localhost:8080")

}
