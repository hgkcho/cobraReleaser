package main

import (
	"fmt"
	"sync"
)

func main() {
	// ch := make(chan string)
	wg := new(sync.WaitGroup)
	mux := new(sync.Mutex)

	var str = "hello"

	wg.Add(1)
	go func() {
		defer wg.Done()
		mux.Lock()
		str = "hage"
		mux.Unlock()
		// ch <- str
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mux.Lock()
		str = "unko"
		mux.Unlock()
	}()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println(<-ch)
	// }()
	wg.Wait()
	fmt.Println(str)
}
