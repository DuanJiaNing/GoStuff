package app

import (
	"fmt"
	"log"
	"net/http"

	"GoStuff/web-middleware/duan"
	"GoStuff/web-middleware/middleware"
)

func init() {
	Router.Use(middleware.Log)
	Router.Use(middleware.Cost)
}

var (
	Router = duan.NewRouter()
)

type RESTfulHandler func(w http.ResponseWriter, r *http.Request) (interface{}, *Error)

func (hp RESTfulHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	model, err := hp(w, r)
	if err != nil {
		log.Printf("ERROR: error when handle request: %v, code: %d, cause %v", err.Message, err.Code, err.Error)
		return
	}

	if _, err := fmt.Fprint(w, model); err != nil {
		log.Printf("ERROR: error when write response: %v", err)
		return
	}
}
