package main

import (
	"fmt"
	"time"
)

func main() {

	do := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("do exit ...")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := do(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling do groutine...")
		close(done)
	}()
	<-terminated
	fmt.Println("Done.")
}
