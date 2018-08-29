package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"context"
	"log"
	"encoding/gob"
	"bytes"
)

func main() {
	r := chi.NewRouter()
	r.Use(MyMiddleWare)
	r.Get("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world\n"))
		writer.Write([]byte("hello world\n"))
		a := request.Context().Value("user")
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(a)
		if err != nil {
			log.Fatal(err)
		}
		//writer.Write([]byte(buf.Bytes()))
		log.Println(a)
	})
	http.ListenAndServe(":3000", r)
}


func MyMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		log.Println("test")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}