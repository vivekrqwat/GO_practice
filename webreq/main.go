package main

import (
	"fmt"
	"net/url"
)

// const url = "https://www.youtube.com"
const url1 = "https://www.google.com/?coursename=golamg&payment=free"

func main() {
	// fmt.Println("WEB REQ")
	// res, er := http.Get(url)
	// if er != nil {
	// 	panic(er)
	// }
	// fmt.Printf("restype %T\n", res)
	// databyte, er := ioutil.ReadAll(res.Body)
	// if er != nil {
	// 	panic(er)
	// }
	// fmt.Println("Data length:", len(databyte))
	// fmt.Println("Data:", string(databyte))
	// defer res.Body.Close()

	res, _ := url.Parse(url1)
	fmt.Println("Parsed URL scheme:", res.Scheme)
	fmt.Println("Parsed Host:", res.Host)
	fmt.Println("Parsed Path:", res.Path)
	fmt.Println("Parsed Port:", res.Port())
	fmt.Println("Parsed Raw Query:", res.RawQuery)

}
