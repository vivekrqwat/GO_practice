package main

import "fmt"

func add(a int, b int) int {
	return (a + b)
}

func name() (string, string, string) {
	return "sam", "vivek", "james"
}
func sqrsum(add func(a int, b int) int) int {
	return add(2, 5) * add(2, 5)
}

func main() {
	fmt.Println("hello", add(5, 10))
	fmt.Println("squr sum", sqrsum(add))

	_, j, s := name()
	fmt.Println(j, s)

}
