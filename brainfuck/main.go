package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	if len(os.Args) == 1 {
		ap.PutRune('\n')
		return
	}

	s := os.Args[1]
	ptr := make([]byte, 1)
	ptr[0] = 0
	p := 0
	i := 0
	for i < len(s) {
		if s[i] == ']' {
			if ptr[p] == 0 {
				i++
				continue
			}
			cnt := 1
			found := false
			for j := i - 1; j >= 0; j-- {
				if s[j] == ']' {
					cnt++
				} else if s[j] == '[' {
					cnt--
					if cnt == 0 {
						i = j
						found = true
						break
					}
				}
			}
			if !found {
				i++
			}
		} else if s[i] == '[' {
			if ptr[p] != 0 {
				i++
				continue
			}
			cnt := 1
			found := false
			for j := i + 1; j < len(s); j++ {
				if s[j] == '[' {
					cnt++
				} else if s[j] == ']' {
					cnt--
					if cnt == 0 {
						i = j
						found = true
						break
					}
				}
			}
			if !found {
				i++
			}
		} else if s[i] == '>' {
			p++
			if p == len(ptr) {
				ptr = append(ptr, byte(0))
			}
			i++
		} else if s[i] == '<' {
			i++
			p--
		} else if s[i] == '+' {
			i++
			ptr[p]++
		} else if s[i] == '-' {
			i++
			ptr[p]--
		} else if s[i] == '.' {
			ap.PutRune(rune(ptr[p]))
			i++
		} else {
			i++
		}
	}
}
