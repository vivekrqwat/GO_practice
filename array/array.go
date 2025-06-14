package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3}
	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(len(nums))

}
