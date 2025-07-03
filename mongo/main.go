package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vivekrqwat/monogo/route"
)

func main() {
	fmt.Println("hello")
	log.Fatal(http.ListenAndServe(":8000", route.Router()))
	fmt.Println("listing to port")
}
