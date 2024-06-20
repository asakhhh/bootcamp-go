package main

import "fmt"

// ZDAT' ZA AZATA!!!!!!!!!!!!

func Split(s string, sep string) []string {
	on_sep := 0
	var t string
	var res []string

	if sep == "" {
		for _, c := range s {
			res = append(res, string(c))
		}
		return res
	}

	for _, c := range s {
		t += string(c)
		if rune(c) == rune(sep[on_sep]) {
			on_sep++
		}
		if on_sep == len(sep) {
			res = append(res, t[:len(t)-len(sep)])
			t = ""
			on_sep = 0
		}
	}
	if len(t) > 0 || string(s[len(s)-1]) == sep {
		res = append(res, t)
	}
	if len(res) == 0 {
		res = append(res, "")
	}

	return res
}

func SquareMatch(c rune, pattern string) bool {
	t := Split(pattern, "-")
	if len(t) == 1 {
		return Contains(pattern, c)
	} else {
		start, end := rune(t[0][0]), rune(t[1][0])
		for start <= end {
			if c == start {
				return true
			}
			start++
		}
		return false
	}
}

func Contains(s string, c rune) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}

func Glob(s, pattern string) bool {
	if len(s) == 0 {
		if len(pattern) == 0 {
			return true
		}
		if len(pattern) == 1 && pattern[0] == '*' {
			return true
		}
		return false
	}
	if len(pattern) == 0 {
		return false
	}

	pts := 0

	for i := range pattern {
		if pts >= len(s) && !(i == len(pattern)-1 && pattern[i] == '*') {
			return false
		}
		if pattern[i] == '[' {
			var cind int
			for j := i + 1; j < len(pattern); j++ {
				if pattern[j] == ']' {
					cind = j
					break
				}
			}
			if SquareMatch(rune(s[pts]), pattern[i+1:cind]) {
				pts++
			} else {
				return false
			}
		} else if pattern[i] == '*' {
			boolean := false
			for j := pts; j < len(s); j++ {
				boolean = boolean || Glob(s[j:], pattern[i+1:])
			}
			return boolean
		} else if pattern[i] == '?' {
			pts++
		} else {
			if s[pts] != pattern[i] {
				return false
			} else {
				pts++
			}
		}
	}
	return pts == len(s)
}

func main() {
	paths := []string{
		"./ ",
		"./run.exe",
		"./README.md",
		"./cmd/main.go",
	}

	patterns := []string{
		"*",
		"*.go",
		"*.?o",
		"*n.[a-z]*",
		"[*]",
	}

	for _, pattern := range patterns {
		fmt.Printf("pattern: %q\n", pattern)
		for _, path := range paths {
			if !Glob(path, pattern) {
				continue
			}

			fmt.Printf("%q\n", path)
		}
		fmt.Println()
	}
}
