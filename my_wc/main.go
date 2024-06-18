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

func Contains(s string, c rune) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}

func Split(s string, sep string) []string {
	on_sep := 0
	var t string
	var res []string

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
	if len(t) > 0 {
		res = append(res, t)
	}
	if len(res) == 0 {
		res = append(res, "")
	}

	return res
}

func ToString(n int) string {
	var t string

	if n < 0 {
		t += "-"
		n = -n
	}

	p, l := 1, 1
	for ; p*10 <= n; p *= 10 {
		l++
	}

	for i := 0; i < l; i++ {
		t += string(rune('0' + (n/p)%10))
		p /= 10
	}
	return t
}

func Join(elements []int, sep string) string {
	var t string
	for i, v := range elements {
		t += ToString(v)
		if i != len(elements)-1 {
			t += sep
		}
	}
	return t
}

func LineCount(s string) int {
	t := Split(s, "\n")
	if t[len(t)-1] == "" {
		return len(t) - 1
	}
	return len(t)
}

func IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func WordCount(s string) int {
	var t string
	ans := 0

	for _, c := range s {
		if IsAlpha(c) {
			t += string(c)
		}
		if (c == ' ' || c == '\n') && t != "" {
			ans++
			t = ""
		}
	}
	if len(t) > 0 {
		ans++
	}

	return ans
}

func Copyof(a []string) []string {
	res := make([]string, len(a))
	copy(res, a)
	return res
}

func main() {
	t := os.Args[1:]
	if len(t) == 0 {
		PrintString("usage: my_wc [-lwc] file\n")
	} else {
		txt, err := os.ReadFile(t[len(t)-1])
		text := string(txt)

		if err != nil {
			PrintString("my_wc: " + t[len(t)-1] + ": No such file or directory\n")
		} else {
			opts := Copyof(t[:len(t)-1])

			if len(opts) == 0 {
				opts = append(opts, "-lwc")
			}
			var res []int
			if Contains(opts[0], 'l') {
				res = append(res, LineCount(text))
			}
			if Contains(opts[0], 'w') {
				res = append(res, WordCount(text))
			}
			if Contains(opts[0], 'c') {
				res = append(res, len(text))
			}

			PrintString(Join(res, " "))

			PrintString(" " + t[0])
			ap.PutRune('\n')
		}
	}
}
