package middleware

import (
	"log"
	"net/http"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("mw in log start")
		next.ServeHTTP(w, r)
		log.Println("mw in log end")
	})
}
