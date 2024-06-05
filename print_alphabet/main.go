package main

import "github.com/alem-platform/ap"

func main() {
	for i := 0; i < 26; i++ {
		ap.PutRune(rune('a' + i))
	}
	ap.PutRune('\n')
}
