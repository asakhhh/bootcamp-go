package main

import "fmt"

func ToLower(s string) string {
	var t string
	ch := 'A' - 'a'
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			t = t + string(v+ch)
		} else {
			t = t + string(v)
		}
	}

	return t
}

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
	if len(t) > 0 {
		res = append(res, t)
	}
	if len(res) == 0 {
		res = append(res, "")
	}

	return res
}

func TextToMorse(s string) string {
	m := map[rune]string{
		'A': ".-",
		'B': "-...",
		'C': "-.-.",
		'D': "-..",
		'E': ".",
		'F': "..-.",
		'G': "--.",
		'H': "....",
		'I': "..",
		'J': ".---",
		'K': "-.-",
		'L': ".-..",
		'M': "--",
		'N': "-.",
		'O': "---",
		'P': ".--.",
		'Q': "--.-",
		'R': ".-.",
		'S': "...",
		'T': "-",
		'U': "..-",
		'V': "...-",
		'W': ".--",
		'X': "-..-",
		'Y': "-.--",
		'Z': "--..",
		'1': ".----",
		'2': "..---",
		'3': "...--",
		'4': "....-",
		'5': ".....",
		'6': "-....",
		'7': "--...",
		'8': "---..",
		'9': "----.",
		'0': "-----",
		'.': ".-.-.-",
		',': "--..--",
		'?': "..--..",
	}
	s = ToLower(s)
	var res string
	for i, c := range s {
		if c == ' ' {
			continue
		}
		res += m[c]
		if i != len(s)-1 {
			res += " "
		}
	}

	return res
}

func MorseToText(s string) string {
	m := map[rune]string{
		'A': ".-",
		'B': "-...",
		'C': "-.-.",
		'D': "-..",
		'E': ".",
		'F': "..-.",
		'G': "--.",
		'H': "....",
		'I': "..",
		'J': ".---",
		'K': "-.-",
		'L': ".-..",
		'M': "--",
		'N': "-.",
		'O': "---",
		'P': ".--.",
		'Q': "--.-",
		'R': ".-.",
		'S': "...",
		'T': "-",
		'U': "..-",
		'V': "...-",
		'W': ".--",
		'X': "-..-",
		'Y': "-.--",
		'Z': "--..",
		'1': ".----",
		'2': "..---",
		'3': "...--",
		'4': "....-",
		'5': ".....",
		'6': "-....",
		'7': "--...",
		'8': "---..",
		'9': "----.",
		'0': "-----",
		'.': ".-.-.-",
		',': "--..--",
		'?': "..--..",
	}

	mp := make(map[string]rune)

	for i, v := range m {
		mp[v] = i
	}
	var res string
	t := Split(s, " ")
	for _, v := range t {
		res += string(mp[v])
	}

	return res
}

func main() {
	fmt.Println(MorseToText("... --- ..."))                                                       // SOS
	fmt.Println(MorseToText("... .- .-.. . -- --..-- .... --- .-- .- .-. . -.-- --- ..- ..--..")) // SALEM,HOWAREYOU?
}
