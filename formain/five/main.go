package main

import (
	"fmt"
	"time"
)

func main() {
	flushed := make(chan struct{})
	go func() {
		defer func() {
			close(flushed)
			fmt.Println("chanel closed")
		}()
		anth()
	}()
	flushed <- struct{}{} // Will panic.
	<-flushed
	fmt.Println("closed")
}

func anth() {
	fmt.Println("sleeping...")
	time.Sleep(4 * time.Second)
}
