package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//http.Handle("/", http.FileServer(http.Dir("c:/")))
	//http.ListenAndServe(":8080", nil)
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("a")))
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("b")))

}
