package main

import "fmt"

func main() {
	fmt.Println("inside main todo app")
	todo := Todos{}
	storage := newStrogare[Todos]("todo.json")
	storage.Load(&todo)
	command := Newcomd()
	command.Execute(&todo)
	storage.Save(todo)

}
