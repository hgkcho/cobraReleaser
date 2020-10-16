package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	btn := Button{Clicked: sync.NewCond(new(sync.Mutex))}
	subscribe := func(c *sync.Cond, fn func()) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		wg.Wait()
	}

	var clickRegisterd sync.WaitGroup
	clickRegisterd.Add(3)
	subscribe(btn.Clicked, func() {
		fmt.Println("Maximizing window")
		clickRegisterd.Done()
	})
	subscribe(btn.Clicked, func() {
		fmt.Println("display annoying dialog box")
		clickRegisterd.Done()
	})
	subscribe(btn.Clicked, func() {
		fmt.Println("mouse cliked")
		clickRegisterd.Done()
	})
	btn.Clicked.Broadcast()
	clickRegisterd.Wait()

}
