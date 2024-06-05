package main

import "github.com/alem-platform/ap"

func main() {
	for i := 0; i < 10; i++ {
		ap.PutRune(rune('0' + i))
	}
	ap.PutRune('\n')
}
