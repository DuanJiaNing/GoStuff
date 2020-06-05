package main

import (
	"bytes"
	"fmt"
	"io"
)

type str struct {
	s        []byte
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func main() {

	var s *str
	if s == nil {
		fmt.Println("1")
	}

	var p *bytes.Reader
	if p != nil {
		fmt.Println("2")
	}
	c(p)
	c(nil)
}

func c(b io.Reader) {
	if b != nil {
		fmt.Println("3")
	}
}
