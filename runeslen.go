package bootcamp

import "github.com/alem-platform/ap"

func RunesLen(arr [20]rune) int {
	c := 0
	
	for c < 20 && arr[c] != '\0' {
		c++
	}

	return c
}