package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func PrintString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func main() {
	for _, v := range os.Args[1:] {
		PrintString(v)
		ap.PutRune('\n')
	}
}
