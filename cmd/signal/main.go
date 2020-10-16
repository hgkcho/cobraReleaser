package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("Done.")
}

// func main() {
//     sig := make(chan os.Signal, 1)
//     signal.Notify(sig,
//         syscall.SIGKILL,
//         syscall.SIGTERM,
//         syscall.SIGINT,
//         os.Interrupt)

//     s := <- sig

//     fmt.Println("signal: ", s)
// }
