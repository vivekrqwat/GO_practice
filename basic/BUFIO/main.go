package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// type User struct {
// 	name string
// 	job  *job
// }
// type job struct {
// 	salary int
// }

type Student struct {
	Name   string `json:"naam"`
	rollno int
}

//	func (u *User) n() {
//		fmt.Println("in func", u)
//		u.name = "james"
//	}
func main() {
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// fmt.Println("reading value", reader.Size(), input)
	// name := "sam rawat"
	// ans := strings.TrimSpace(name)
	// fmt.Printf("%T", ans)
	// k := time.Now()
	// mynum := 44
	// var p = &mynum
	// fmt.Printf("time is%s", k.Format("2006-01-02"))

	// fmt.Println("", *p)
	// mp := make(map[int]int, 6)
	// mp[0] = 6
	// fmt.Println("map value is", mp)
	// for key, value := range mp {
	// 	fmt.Println(key, value)
	// }
	// delete(mp, 0)
	// fmt.Println("map value is", mp)
	// u := User{
	// 	name: "Sam",
	// 	job:  &job{salary: 1000},
	// }
	// u.n()
	// fmt.Println(u.name)

	// file, er := os.Create("./myfile.txt")
	// if er != nil {
	// 	panic(er)
	// }
	// len, er := io.WriteString(file, "sam was here")

	// if er != nil {
	// 	panic(er)
	// }
	// fmt.Println(len)
	// readfile("myfile.txt")
	// defer file.Close()

	data := []Student{
		{"vivek", 78},
		{"kamal", 73},
	}

	encode(data)
	//decode
	jsondata := []byte(`[{"naam":"vivek"},{"naam":"kamal"}]`)
	var st []Student
	check := json.Valid(jsondata)

	if check {
		json.Unmarshal(jsondata, &st)
		fmt.Println("", st)
	} else {
		fmt.Println("no valid")
	}
}

func readfile(filename string) {
	data, er := os.ReadFile(filename)
	if er != nil {
		panic(er)
	}
	fmt.Println("", string(data))
}
func encode(data []Student) {
	value, er := json.Marshal(data)
	if er != nil {
		panic(er)
	}
	fmt.Printf("json value: %s\n", string(value))

}
