package main

import "fmt"

func main() {
	// var nums []int
	var nums = make([]int, 5)

	// nums = append(nums, 1)
	// nums := []int{}
	fmt.Println(nums)
	fmt.Println(nums[0:2])
	fmt.Println("length", len(nums))

}
