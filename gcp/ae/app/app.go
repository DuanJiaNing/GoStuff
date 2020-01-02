package app

import (
	"GoStuff/gcp/ae/log"
	"net/http"
)

func Handler(handle func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := log.EmbedToContext(r)
		r = r.WithContext(ctx)
		handle(w, r)

		log.Flush(ctx)
	}
}
