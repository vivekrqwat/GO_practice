package main

import (
	"fmt"
	"sync"
)

func p1(i int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("p1", i)
}
func p2(i int) {
	fmt.Println("p2", i)
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 5; i++ {
		w.Add(1)
		go p1(i, &w)
		// go p2()

	}
	w.Wait()
	// time.Sleep(time.Second * 2)

}
