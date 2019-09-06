package main

import (
	"net/http"

	"github.com/deliangyang/hello-test/middleware"
	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Verify)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe("127.0.0.1:4000", r)
}
