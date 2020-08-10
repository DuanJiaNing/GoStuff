package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test/{type:(?:ac|b123)}", test)
	//http.HandleFunc("/fu", fileUpload)
	http.Handle("/", r)

	if err := http.ListenAndServe(":1313", nil); err != nil {
		log.Fatal(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("www")
}

func fileUpload(w http.ResponseWriter, r *http.Request) {
	log.Println("got file")

	mr, err := r.MultipartReader()
	if err != nil {
		printError(err)
		return
	}

	for {
		part, err := mr.NextPart()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				printError(err)
				return
			}
		}

		all, err := ioutil.ReadAll(part)
		if part.FormName() == "text_a" {
			log.Println(part.FormName(), part.FileName(), len(all), string(all), err)
		} else {
			log.Println(part.FormName(), part.FileName(), len(all), err)
		}
	}
}

func printError(err error) {
	log.Printf("err:%v", err)
}
