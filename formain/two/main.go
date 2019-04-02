package main

import (
	"fmt"
	"github.com/gorilla/mux"
	. "net/http"
	"os"
)

type OptionsHandler struct {
}

type MyHandler struct {
}

type QuitHandler struct {
}

func (h *QuitHandler) ServeHTTP(w ResponseWriter, r *Request) {
	fmt.Println("server dead")
	os.Exit(1)
}

func (h *OptionsHandler) ServeHTTP(w ResponseWriter, r *Request) {
	fmt.Printf("OptionsHandler: %s : %s\n", r.Method, r.URL)
}

func (h *MyHandler) ServeHTTP(w ResponseWriter, r *Request) {
	fmt.Printf("MyHandler: %s : %s\n", r.Method, r.URL)
}

var Router = mux.NewRouter()

func main() {

	optionsHandler := &OptionsHandler{}
	Router.PathPrefix("/api").Methods(MethodOptions).Handler(optionsHandler)

	myHandler := &MyHandler{}
	subRouter := Router.PathPrefix("/api").Subrouter()
	subRouter.Path("/quit").Methods(MethodGet).Handler(&QuitHandler{})
	subRouter.Path("/m").Methods(MethodGet).Handler(myHandler)
	subRouter.Path("/b").Methods(MethodGet).Handler(myHandler)
	subRouter.Path("/a").Methods(MethodGet).Handler(myHandler)
	subRouter.Path("/a").Methods(MethodOptions).Handler(myHandler) // OPTIONS request still handed by OptionsHandler

	srv := &Server{
		Handler: Router,
		Addr:    "localhost:8080",
	}

	fmt.Println("server started...")
	srv.ListenAndServe()

}
