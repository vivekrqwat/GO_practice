package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todoList = []todo{
	{ID: "1", Name: "vivek", Completed: true},
	{ID: "2", Name: "sachin", Completed: false},
}

func getTodod(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todoList)
}
func addtodo(c *gin.Context) {
	var newtodo todo
	if er := c.BindJSON(&newtodo); er != nil {
		return
	}
	todoList = append(todoList, newtodo)
	c.IndentedJSON(http.StatusCreated, newtodo)
}

func gettodobyid(id string) (*todo, error) {
	for i, t := range todoList {
		if t.ID == id {
			return &todoList[i], nil
		}
	}
	return nil, errors.New("todo not found")

}

func updatetodo(c *gin.Context) {
	id := c.Param("id")
	// fmt.Printf("%T\n", id)
	todo, er := gettodobyid(id)
	if er != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func updatetoggle(c *gin.Context) {
	id := c.Param("id")
	todo, er := gettodobyid(id)
	if er != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	c.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todo", getTodod)
	router.POST("/todo", addtodo)
	router.GET("/todo/:id", updatetodo)
	router.PATCH("/todo/:id", updatetoggle)
	router.Run("localhost:8080")

}
