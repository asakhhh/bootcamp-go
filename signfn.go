package main

import "github.com/alem-platform/ap"

func Sign(n int) {
	if n > 0 {
		ap.PutRune('+')
	} else if n == 0 {
		ap.PutRune('0')
	} else {
		ap.PutRune('-')
	}
	return
}

func main() {
	Sign(2024)
}
