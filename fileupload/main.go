package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/fu", fileUpload)
	if err := http.ListenAndServe(":1313", nil); err != nil {
		log.Fatal(err)
	}
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

		log.Println(part.FormName(), part.FileName())
	}
}

func printError(err error) {
	log.Printf("err:%v", err)
}
