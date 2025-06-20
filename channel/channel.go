package main

import (
	"fmt"
	"time"
)

//	func call(a chan int) {
//		fmt.Println("we get", <-a)
//	}

// recv
// func add(res chan int, a int, b int) {
// 	numres := a + b
// 	res <- numres

// }

func callemail(e chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range e {
		fmt.Println("your email is", email)
		time.Sleep(time.Second)
	}
}

func main() {
	// msg := make(chan int)

	// go call(msg)
	// msg <- 5                    //callin it before go routine will lead deadloack
	// time.Sleep(time.Second * 2) //to stop main function so tha it wait till call get execute

	// msg := make(chan string)
	// msg <- "sam"
	// rec := <-msg
	// fmt.Println("rec", rec)

	// recv

	// num := make(chan int)
	// go add(num, 5, 5)
	// res := <-num
	// fmt.Println("result we get", res)

	//
	// done := make(chan bool)
	// emailchan := make(chan string, 100)
	// go callemail(emailchan, done)

	// for i := 0; i < 10; i++ {
	// 	emailchan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// close(emailchan) // close after all sends
	// fmt.Println("done processing")
	// <-done //to stop the main function before execution of callemail

	//
	chan1 := make(chan string)
	chan2 := make(chan int)
	go func() {
		chan1 <- "sam"
	}()
	go func() {
		chan2 <- 1
	}()
	for i := 0; i < 2; i++ { //need to use for loop  to catch all channel
		select {
		case ch1 := <-chan1:
			{
				fmt.Println("chna11", ch1)
			}

		case ch2 := <-chan2:
			{
				fmt.Println("chan2", ch2)
			}
		}
	}

}
