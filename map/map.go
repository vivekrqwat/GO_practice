package main

import (
	"fmt"
	"maps"
)

func main() {

	// m := make(map[string]string)
	m := map[string]string{"vivek": "hero2"}
	m2 := map[string]string{"vivek": "hero2"}
	//setting
	m["sam"] = "hero"
	//if key is not present in map it return zero value
	delete(m, "sam")
	v, ok := m["vivek"]
	fmt.Println(m["sam"], m["vivek"])
	if ok {
		fmt.Println("all ok", v, ok)
	} else {
		fmt.Println("not ok")
	}
	fmt.Println(maps.Equal(m, m2))

}
