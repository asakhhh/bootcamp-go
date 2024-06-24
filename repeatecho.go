package bootcamp

import "fmt"

var match map[int]int

func RepeatEcho(s string) string {
	var br []int

	match = make(map[int]int)

	for i, c := range s {
		if c == '(' {
			br = append(br, i)
		} else if c == ')' {
			if len(br) != 0 {
				op := br[len(br)-1]
				match[op] = i
				br = br[:len(br)-1]
			}
		}
	}

	// return ""
	return RepeatEcho1(s, 1, 0, len(s)-1)
}

func GetNumber(s string, ind int) int {
	if ind < 0 || s[ind] < '0' || s[ind] > '9' {
		return -1
	}
	num := 0
	p := 1
	for ind >= 0 && s[ind] >= '0' && s[ind] <= '9' {
		num += int(s[ind]-'0') * p
		ind--
		p *= 10
	}

	return num
}

func RepeatEcho1(s string, cnt, l, r int) string {
	var res string
	i := l

	fmt.Println(l, r, cnt)

	for i <= r {
		if s[i] == '(' {
			closing, found := match[i]
			rep := GetNumber(s, i-1)
			if !found || rep == -1 {
				res += "("
				i++
			} else {
				res += RepeatEcho1(s, rep, i+1, closing-1)
				i = closing + 1
			}
		} else {
			res += string(s[i])
			i++
		}
	}

	var t string
	for cnt > 0 {
		t += res
		cnt--
	}

	return t
}

func main() {
	//                      012345678901234567890123
	fmt.Println(RepeatEcho("Ba2(na), 2(d2(a)) 10(!)(")) // Banana daadaa !!!!!!!!!!(
}
