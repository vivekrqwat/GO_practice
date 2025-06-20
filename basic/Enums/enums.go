package main

import "fmt"

// type mystring string

// func (a mystring) sam() {
// 	fmt.Println(a)
// }

type orderstatus float32

const (
	Recived orderstatus = iota
	confirmed
	Prepared
	Delivered
)

func changestatus(status orderstatus) {
	fmt.Println("change status", status)
}
func main() {
	// var a mystring = "he;p"
	// a.sam()
	changestatus(Delivered)
	fmt.Println("enums")

}
