package main

import (
	"GoStuff/web-middleware/app"
	"log"
	"net/http"

	_ "GoStuff/web-middleware/handler"
)

func main() {
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", app.Router); err != nil {
		log.Fatal(err)
	}
}
