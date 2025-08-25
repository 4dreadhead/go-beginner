package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8081", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}
