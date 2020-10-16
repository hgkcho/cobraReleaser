package main

import "fmt"

func mani() {
	gen := generator()
	gen()
	gen()
	gen()
}

func generator() func() int {
	x := 1
	return func() int {
		x++
		fmt.Println(x)
		return x
	}
}
