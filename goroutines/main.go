package main

import (
	"fmt"
	"sync"
)

var a = []string{
	"https://github.com",
	"https://fb.com",
	"https://google.com",
}
var w sync.WaitGroup
var m sync.Mutex
var g = []string{"test"}

func main() {
	// fmt.Println("goroutine")
	// go greet("Sam")
	// go greet("vivek")
	// time.Sleep(3 * time.Millisecond)

	// for _, web := range a {
	// 	w.Add(1)
	// 	go getStatuscode(web)

	// }

	// w.Wait()
	// for _, r := range g {
	// 	w.Add(1)
	// 	fmt.Println(",", r)

	// }

	//channel use for communication between go routines ,they are block statement
	wg := &sync.WaitGroup{}
	ch := make(chan int, 2)
	wg.Add(2)
	go func(wg *sync.WaitGroup, ch chan int) {
		defer wg.Done()
		j, er := <-ch

		fmt.Println(j)
		fmt.Println("received", er)

	}(wg, ch)

	go func(wg *sync.WaitGroup, ch chan int) {
		defer wg.Done()
		fmt.Println("sending...")
		close(ch)
		ch <- 5
		ch <- 10
	}(wg, ch)

	wg.Wait()
}

// func greet(s string) {
// 	for i := 0; i < 6; i++ {

// 		fmt.Println(s, "...", i)

//		}
//	}
// func getStatuscode(s string) {
// 	defer w.Done()
// 	res, err := http.Get(s)
// 	if err != nil {
// 		fmt.Println("error")
// 		return
// 	}
// 	m.Lock()
// 	g = append(g, s)
// 	m.Unlock()
// 	fmt.Printf("%dstatus code%s\n", res.StatusCode, s)

// }
