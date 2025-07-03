package main

import "fmt"

// type paymenter interface {
// 	pay(amount int)
// }

// type payment struct {
// 	gateway paymenter
// }

// func (p payment) transfer(amount int) {
// 	// p1 := payment
// 	p.gateway.pay(amount)
// }

// type razorpy struct{}

// func (r razorpy) pay(amount int) {
// 	fmt.Println("razor pay amount", amount)
// }

type g interface {
	check(s string)
}
type student struct{}

func (s *student) check(s1 string) {
	fmt.Println(s1, "is student")
}

type teacher struct{}

func (t *teacher) check(t1 string) {
	fmt.Println(t1, "is teacher")
}

type final struct {
	name g
}

func (f *final) finalcheck() {
	f.name.check("sam")
}
func main() {
	// r1 := razorpy{}
	// p1 := payment{
	// 	gateway: r1,
	// }
	// p1.transfer(100)
	st := final{
		name: &teacher{},
	}
	st.finalcheck()

}
