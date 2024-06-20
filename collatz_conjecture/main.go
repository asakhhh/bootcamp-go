package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func toint(s string) int {
	ans := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		ans = ans*10 + int(c-'0')
	}
	return ans
}

func PrintNum(n int) {
	p := 1
	for p*10 <= n {
		p *= 10
	}
	for p > 0 {
		ap.PutRune(rune('0' + (n/p)%10))
		p /= 10
	}
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		return
	}
	n := toint(os.Args[1])
	if n < 1 {
		return
	}
	for n > 1 {
		PrintNum(n)
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		ap.PutRune('\n')
	}
	PrintNum(n)
	ap.PutRune('\n')
}
