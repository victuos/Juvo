package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fname := "/helpboard.txt"
	fh, err := os.Open(os.Getenv("HOME") + fname)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, fh)
	if err != nil {
		log.Fatal(err)
	}
}
