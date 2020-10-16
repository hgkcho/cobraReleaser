package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(new(sync.Mutex))
	queue := make([]interface{}, 0, 10)

	addQuere := func(i int) {
		fmt.Fprintf(os.Stdout, "Addinng to queue \n")
		c.L.Lock()
		defer c.L.Unlock()
		queue = append(queue, struct{}{})
		c.Wait()
		fmt.Printf("go %v\n", i)
	}

	for i := 0; i < 10; i++ {
		go addQuere(i)
	}

	for i := 3; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	c.Broadcast()
	time.Sleep(2 * time.Second)

}
