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
	t := os.Args[1:]
	if len(t) == 0 {
		PrintString("usage: my_cat file ...")
	} else {
		for _, v := range t {
			txt, err := os.ReadFile(v)
			if err != nil {
				PrintString("my_cat: " + v + ": No such file or directory")
			} else {
				PrintString(string(txt))
			}
			ap.PutRune('\n')
		}
	}
}
