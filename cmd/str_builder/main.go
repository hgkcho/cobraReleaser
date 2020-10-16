package main

import (
	"fmt"
	"strings"
)

func main() {
	var buf strings.Builder
	buf.WriteString("hoge")
	buf.Write([]byte(":"))
	fmt.Println(buf.String())
}
