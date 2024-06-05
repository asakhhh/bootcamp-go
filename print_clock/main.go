package main

import "github.com/alem-platform/ap"

func main() {
	hh := 0
	mm := 0

	for t := 0; t < 1440; t++ {
		ap.PutRune(rune('0' + hh/10))
		ap.PutRune(rune('0' + hh%10))
		ap.PutRune(':')
		ap.PutRune(rune('0' + mm/10))
		ap.PutRune(rune('0' + mm%10))
		ap.PutRune('\n')
		mm++
		if mm == 60 {
			mm = 0
			hh++
		}
	}
}
