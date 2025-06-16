package main

import (
	"fmt"
	"time"
)

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o *order) changestatus(status string) {
	o.status = status
}

func newo(id1 string, amount1 float32, status1 string) *order {
	o := order{
		id:     id1,
		amount: amount1,
		status: status1,
	}
	return &o
}

func main() {
	var o order = order{
		id:     "1",
		amount: 50,
		status: "received",
	}
	l := struct {
		name string
	}{"sam"}

	var k *order = newo("1", 50, "same")

	o.changestatus("go it")
	o.createdAt = time.Now()
	fmt.Println("my order", o, *k, l)

}
