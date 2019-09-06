package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})
	http.ListenAndServe(":9090", nil)
}
