package bootcamp

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

// func main() {
// 	fmt.Println(TextToMorse("SOS"))                 // ... --- ...
// 	fmt.Println(TextToMorse("salem, how are you?")) // ... .- .-.. . -- --..-- .... --- .-- .- .-. . -.-- --- ..- ..--..
// }
