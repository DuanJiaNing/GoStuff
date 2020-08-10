package main

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type New string

func (n New) Error() string {
	return string(n)
}

func main() {
	//
	//err := errors.New("a")
	//err = New("b")

	fmt.Println(runtime.NumCPU())
	//fmt.Fprintf()
	strings.Split("aaaa", "a")
	errors.New("")
}
