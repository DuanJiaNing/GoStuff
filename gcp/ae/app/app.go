package app

import (
	"GoStuff/gcp/ae/log"
	"context"
	"net/http"
)

func Handler(handle func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(newContext(r))
		ctx := r.Context()
		handle(w, r)

		log.Flush(ctx)
	}
}

func newContext(r *http.Request) context.Context {
	return log.NewContext(r)
}
