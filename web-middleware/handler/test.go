package handler

import (
	"GoStuff/web-middleware/app"
	"net/http"
	"time"
)

func init() {
	app.Router.Path("/test", app.RESTfulHandler(test))
}

func test(w http.ResponseWriter, r *http.Request) (interface{}, *app.Error) {
	time.Sleep(time.Second * 3)
	return "reach me", nil
}
