package main

import (
	"net/http"
	"fmt"
)

func main() {

	http.HandleFunc("/hello", test)
	http.ListenAndServe(":8888", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
	var str = "Hello world"
	content := [] byte(str)
	w.Write(content)
}