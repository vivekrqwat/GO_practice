package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Add    string
	del    int
	toggle int
	edit   string
	list   bool
}

func Newcomd() *Command {
	cf := Command{}
	flag.StringVar(&cf.Add, "add", "", "add new todo")
	flag.StringVar(&cf.edit, "edit", "", "edit todo")
	flag.IntVar(&cf.del, "del", -1, "delete todo by index")
	flag.IntVar(&cf.toggle, "toogle", -1, "toogle todo give us index")
	flag.BoolVar(&cf.list, "list", false, "list all todos")
	flag.Parse()
	return &cf
}

func (c *Command) Execute(todo *Todos) {
	switch {
	case c.Add != "":
		todo.add(c.Add)
	case c.del != -1:
		todo.del(c.del)
	case c.list:
		todo.print()
	case c.edit != "":
		parts := strings.SplitN(c.edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid command use index:title")
			os.Exit(1)
		}
		index, er := strconv.Atoi(parts[0])
		if er != nil {
			fmt.Println("invalid index", parts[0])
			os.Exit(1)
		}
		todo.edit(index, parts[1])
	case c.toggle != -1:
		todo.toggle(c.toggle)

	default:
		fmt.Println("invalid command")

	}

}
