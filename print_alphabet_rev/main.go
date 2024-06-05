package main

import "github.com/alem-platform/ap"

func main() {
	for i := 25; i >= 0; i-- {
		ap.PutRune(rune('a' + i))
	}
	ap.PutRune('\n')
}
