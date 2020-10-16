package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/hgkcho/cobraReleaser/pkg/shell"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	command := os.Getenv("SHELL")
	// command := os.Getenv("EDITOR")
	if command == "" {
		command = "sh"
	}
	resp, err := http.Get("https://raw.githubusercontent.com/zdharma/zinit/master/doc/install.sh")
	// resp, err := http.Get("https://raw.githubusercontent.com/hgkcho/sumcp/master/README.md")

	handleErr(err)
	defer resp.Body.Close()

	tmpf, err := ioutil.TempFile("", "tmp")
	handleErr(err)
	defer os.Remove(tmpf.Name())
	defer tmpf.Close()
	// var buf bytes.Buffer
	// resp.Write(&buf)
	// fmt.Println(buf.String())
	resp.Write(tmpf)
	b, _ := ioutil.ReadFile(tmpf.Name())
	fmt.Println(string(b))

	sh := shell.New(command, tmpf.Name())
	err = sh.Run(context.Background())
	handleErr(err)

}
