package main

import "fmt"

type paymenter interface {
	pay(amount int)
}

type payment struct {
	gateway paymenter
}

func (p payment) transfer(amount int) {
	// p1 := payment
	p.gateway.pay(amount)
}

type razorpy struct{}

func (r razorpy) pay(amount int) {
	fmt.Println("razor pay amount", amount)
}

func main() {
	r1 := razorpy{}
	p1 := payment{
		gateway: r1,
	}
	p1.transfer(100)

}
