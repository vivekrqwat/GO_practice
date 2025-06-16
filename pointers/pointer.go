package main

import "fmt"

func changenum(num *int) {
	*num = 6
	fmt.Println("change nums", *num)
}

func main() {
	num := 4
	changenum(&num)
	fmt.Println("nums", num)
}
