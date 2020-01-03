package main

import (
	"GoStuff/web-middleware/app"
	"log"
	"net/http"

	_ "GoStuff/web-middleware/handler"
)

func init() {
	http.Handle("/", app.Router)
}

func main() {
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
