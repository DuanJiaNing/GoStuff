package middleware

import (
	"log"
	"net/http"
	"time"
)

func Cost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("mw in cost start")
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("mw in cost end: %vs\n", time.Now().Sub(start).Seconds())
	})
}
