package main

import "fmt"

//vardic function cam take  n number of arguments
func vard(nums ...interface{}) []interface{} {
	// total := 0
	// for _, num := range nums {
	// 	total += num
	// }
	return nums

}

//clouser
func count() func() int {
	var c int = 0
	return func() int {
		c += 1
		return c
	}
}

func main() {
	c := count()

	fmt.Println("fk", vard(1, 2, 3, "sam", 5))
	fmt.Println("first count", c())
	fmt.Println("second count", c())
}
