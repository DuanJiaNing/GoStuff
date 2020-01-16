package main

import (
	"io"
	"io/ioutil"
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
