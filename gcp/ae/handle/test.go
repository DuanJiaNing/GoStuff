package handle

import (
	"net/http"

	"GoStuff/gcp/ae/app"
	"GoStuff/gcp/ae/log"
)

func init() {
	http.HandleFunc("/api/test", app.Handler(test))
}

func test(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Debugf(ctx, "log test debug")
	log.Infof(ctx, "log test info")
	log.Warningf(ctx, "log test warning")
}
