package main

import "fmt"

func TestNthRune(fn func(s string, n int) rune) bool {
	s := "\n abacabac11\n"

	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if fn(s, i+1) != r[i] {
			return false
		}
	}
	if fn(s, 0) >= 0 || fn(s, len(s)+1) >= 0 {
		return false
	}
	return true
}

func NthRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return -1
	}
	runes := []rune(s)
	return runes[n-1]
}

func ZeroRune(s string, n int) rune {
	return '0'
}

func FirstRune(s string, n int) rune {
	if len(s) > 0 {
		return rune(s[0])
	}
	return -1
}

func main() {
	fmt.Println(TestNthRune(NthRune))   // true
	fmt.Println(TestNthRune(ZeroRune))  // false
	fmt.Println(TestNthRune(FirstRune)) // false
}
