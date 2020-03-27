package main

import (
	"fmt"
	"time"
)

func main() {
	stopFlushing := make(chan int, 1)
	go doth(stopFlushing)

	anth()
	stopFlushing <- 1
}

func anth() {
	fmt.Println("sleeping...")
	time.Sleep(4 * time.Second)
}

func doth(stop <-chan int) {
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-stop:
			// Request finished.
			tick.Stop()
			return
		case <-tick.C:
			fmt.Println("doth")
		}
	}
}
