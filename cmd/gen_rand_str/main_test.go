package main

import (
	"bytes"
	"math/rand"
	"testing"
)

func Benchmark_byte_init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		length := 10
		lenChar := len(CharAll)

		buf := make([]byte, lenChar)

		for i := 0; i < length; i++ {
			char := CharAll[rand.Intn(lenChar)]
			buf = append(buf, char)
		}
	}
}

func Benchmark_byte_init2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		length := 10
		lenChar := len(CharAll)

		buf := make([]byte, length)

		for i := 0; i < length; i++ {
			buf = append(buf, CharAll[rand.Intn(lenChar)])
		}
	}
}

func Benchmark_byte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var m2 = bytes.NewBuffer(make([]byte, 0, 100))
		length := 10
		lenChar := len(CharAll)

		for i := 0; i < length; i++ {
			_ = generateTwoChar(lenChar)
		}
	}
}

func Benchmark_buffers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// var m2 = bytes.NewBuffer(make([]byte, 0, 100))
		pw := &bytes.Buffer{}
		length := 10
		lenChar := len(CharAll)

		for i := 0; i < length; i++ {
			char := CharAll[rand.Intn(lenChar)]
			_ = pw.WriteByte(char)
		}
	}
}
