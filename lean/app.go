package main

import (
	"fmt"
	"time"
)

func main1() {
	//fmt.Println("method from package f1")
	//sayHi("Hi")

	a1()
}

func sayHi(str string) {
	for i := 0; i < 10; i++ {
		sl := 100 * time.Millisecond
		time.Sleep(sl)
		P_("after sleep: ", sl, str)
	}
}

func P_(p ...interface{}) {
	fmt.Println(p)
}

func a1() {
	s := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	P_(x, y, x+y)
}

func sum(arr []int, ch chan int) {
	var sum = 0
	for _, va := range arr {
		sum += va
	}
	ch <- sum
}
