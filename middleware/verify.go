package middleware

import (
	"net/http"
	"strings"
)

func Verify(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-type", "application/json")
		w.Header().Add("Content-Type1", "application/json")
		w.WriteHeader(200)

		w.Write([]byte(r.Header.Get("app-version") + "\n"))
		w.Write([]byte("hello world\n"))
		for k, v := range r.Header {
			w.Write([]byte(k))
			w.Write([]byte(":"))
			w.Write([]byte(strings.Join(v, ",")))
			w.Write([]byte("\n"))
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}