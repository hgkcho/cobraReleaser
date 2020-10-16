package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Password struct {
	Title       string   `json:"title"`
	Account     string   `json:"account"`
	Descripiton string   `json:"description"`
	PasswordSet []string `json:"PasswordSet"`
}

const (
	digits = "0123456789"
	upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower  = "abcdefghijklmnopqrstuvwxyz"
	// syms   = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	syms = "!#$%&()*+,-.<=>?@[]_{}"
	// CharAlpha is the class of letters
	CharAlpha = upper + lower
	// CharAlphaNum is the class of alpha-numeric characters
	CharAlphaNum = digits + upper + lower
	// CharAll is the class of all characters
	CharAll = digits + upper + lower + syms
)

func generateTwoChar(max int) []byte {
	// i, err := crand.Int(crand.Reader, big.NewInt(int64(max)))
	// handleErr(err)

	// if err == nil {
	// 	return int(i.Int64())
	// }
	// fmt.Fprintln(os.Stderr, "WARNING: No crypto/rand available. Falling back to PRNG")
	var ret []byte
	rand.Seed(time.Now().UnixNano())
	ret = append(ret, CharAll[rand.Intn(max)])
	ret = append(ret, CharAll[rand.Intn(max)])
	return ret
}

func main() {
	// g := xkcdpwgen.NewGenerator()
	// g.SetNumWords(2)
	// g.SetCapitalize(true)
	// if err := g.UseLangWordlist("de"); err != nil {
	// 	log.Fatal(err)
	// }
	// password := g.GeneratePasswordString()
	// fmt.Println(password)
	pw := Password{
		Title:       "tmp",
		Account:     "pulpul",
		Descripiton: "hahahaha",
		PasswordSet: []string{},
	}

	length := 5 * 5
	var buf []byte
	// var pwSet = make([]string, length)
	lenChar := len(CharAll)
	ch := make(chan string, length)
	wg := new(sync.WaitGroup)

	for i := 0; i < length; i++ {
		// char := CharAll[generateTwoChar(lenChar)]
		wg.Add(1)
		go func() {
			defer wg.Done()
			buf = generateTwoChar(lenChar)
			ch <- string(buf)
		}()
		// fmt.Println(string(buf))
		// pw.PasswordSet = append(pw.PasswordSet, string(buf))
	}
	wg.Wait()

	for i := 0; i < length; i++ {
		char := <-ch
		pw.PasswordSet = append(pw.PasswordSet, char)
	}

	f, err := os.Create("tmp.txt")
	defer f.Close()
	handleErr(err)

	json.NewEncoder(f).Encode(&pw)
	pw.render()
}

func (p Password) render() {
	for k, v := range p.PasswordSet {
		if k == 0 {
			fmt.Println("------------------------------------")
		}
		if k%5 == 4 {
			fmt.Fprintf(os.Stdout, "|  %s  |\n", v)
			fmt.Println("                                            ")
		} else {
			fmt.Fprintf(os.Stdout, "|  %s  ", v)
		}
		if k == 24 {
			fmt.Println("------------------------------------")
		}
	}
}
