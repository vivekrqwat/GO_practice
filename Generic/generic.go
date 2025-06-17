package main

import "fmt"

type Sorint interface {
	string | int
}

func call[T string | int](items []T) { // can also use compareble which is an nterfsace having all data types
	for _, i := range items {
		fmt.Println("", i)
	}
}

type stack[T int | string] struct {
	mystack []T
}

func main() {
	stack := stack[int]{
		mystack: []int{1, 2, 3, 4},
	}
	fmt.Println(stack)
	lsit := []int{1, 2, 3, 4, 5}
	lsit2 := []string{"sam", "vivek"}
	call(lsit)
	call(lsit2)
}
