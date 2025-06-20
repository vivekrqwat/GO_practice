package main

import (
	"fmt"
	"sync"
)

func (v *post) add(wg *sync.WaitGroup) {
	defer wg.Done()
	v.mu.Lock()
	v.views += 1
	defer v.mu.Unlock()
}

type post struct {
	views int
	mu    sync.Mutex
}

func main() {
	mainpost := post{views: 0}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mainpost.add(&wg)

	}
	wg.Wait()
	fmt.Println("views", mainpost.views)
}
