package main

import (
	"fmt"
	"time"
)

func main() {
	//while
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("i=", i)
	// 	i = i + 1
	// }
	//infinite loop
	// for {
	// 	fmt.Println("sam")
	// }
	//for
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("sam")
	// }
	// for i := range 3 {
	// 	fmt.Println(i)
	// }
	// i := 0
	// if i == 0 || i == 1 {
	// 	fmt.Println("it")
	// }

	//multiple conditon swtich
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("holiday")

	}
	//type
	Woami := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("integer")

		default:
			fmt.Println("other", t)
		}

	}
	Woami("sma")
}
