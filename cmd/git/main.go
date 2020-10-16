package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
	if !errors.Is(err, git.ErrRepositoryAlreadyExists) {
		handleErr(err)
	}

	pwd, err := os.Getwd()
	handleErr(err)

	os.Symlink("/tmp/foo/", filepath.Join(pwd, "go-git"))
	handleErr(err)

}
