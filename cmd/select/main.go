package main

import (
	"fmt"
	"time"
)

func main() {
	// var c1, c2 <-chan interface{}
	// var c3 chan<- interface{}

	// select {
	// case <- c1:
	// 	fmt.Println("c1")
	// case <- c2:
	// 	fmt.Println("c2")
	// case c3 <- struct{}{}:
	// 	fmt.Println("c3")
	// }
	ex()
}

func ex() {
	start := time.Now()
	ch := make(chan interface{})
	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()

	fmt.Println("blocking on read ...")
	select {
	case <-ch:
		fmt.Printf("Unblocked %v later. \n", time.Since(start))
	}
}
