package main

import (
	"os"

	"github.com/alem-platform/ap"
)

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

func stringcmp(s, t string) bool {
	if len(s) < len(t) {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	return s < t
}

func toint(s string) int {
	num := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		num = num*10 + int(c-'0')
	}
	return num
}

func PrintString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func main() {
	if len(os.Args) == 1 {
		PrintString("NV\n")
		return
	}

	arg := os.Args[1]
	if toint(arg) == -1 || stringcmp("2147483647", arg) {
		PrintString("NV")
	} else {
		t := toint(arg)
		d := t / (24 * 60 * 60)
		t %= 24 * 60 * 60
		h := t / 3600
		t %= 3600
		m := t / 60
		s := t % 60
		if d > 0 {
			PrintNum(d)
			ap.PutRune('d')
			ap.PutRune(' ')
		}
		if h > 0 {
			PrintNum(h)
			ap.PutRune('h')
			ap.PutRune(' ')
		}
		if m > 0 {
			PrintNum(m)
			ap.PutRune('m')
			ap.PutRune(' ')
		}
		if s > 0 {
			PrintNum(s)
			ap.PutRune('s')
			ap.PutRune(' ')
		}
	}
	ap.PutRune('\n')
}
