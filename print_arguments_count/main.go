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

func PrintNum(n int) { // to print numbers
	decimal := 1
	for decimal*10 <= n {
		decimal *= 10
	}
	for decimal > 0 {
		ap.PutRune(rune(n/decimal) + '0')
		n = n % decimal
		decimal /= 10
	}
}

func main() {
	PrintNum(len(os.Args) - 1)
	ap.PutRune('\n')
}
