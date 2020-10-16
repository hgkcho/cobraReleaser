package main

import "fmt"

type ByteSize int64

const (
	a = iota
	b
	c
)

const (
	_           = iota // 1番目の値(0)を無視
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Printf("%b\n", KB)
	fmt.Printf("%v\n", KB)
	fmt.Printf("%b\n", MB)
	fmt.Printf("%v\n", MB)
	fmt.Printf("%b\n", GB)
	fmt.Printf("%v\n", GB)
	fmt.Printf("%b\n", TB)
	fmt.Printf("%v\n", TB)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
