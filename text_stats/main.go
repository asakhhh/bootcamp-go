package main

import "fmt"

func ToLower(c rune) rune {
	ch := 'A' - 'a'
	if c >= 'A' && c <= 'Z' {
		return rune(c - ch)
	}
	return c
}

func IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func IsVowel(c rune) bool {
	cc := ToLower(c)
	return cc == 'a' || cc == 'e' || cc == 'i' || cc == 'o' || cc == 'u'
}

func IsConsonant(c rune) bool {
	cc := ToLower(c)
	return IsAlpha(cc) && !IsVowel(cc)
}

func IsNumeric(c rune) bool {
	return c >= '0' && c <= '9'
}

func IsPunct(c rune) bool {
	return c == '.' || c == '!' || c == '?'
}

func IsSpace(c rune) bool {
	return c == ' '
}

func NotEmpty(s string) bool {
	return len(s) > 0
}

func main() {
	var text string
	var r rune

	fmt.Printf("Enter text: ")

	for true {
		fmt.Scanf("%c", &r)
		if r != '\n' {
			text += string(r)
		} else {
			break
		}
	}

	var characters, sentences, words, letters, vowels, consonants, digits, spaces, specials int

	characters = len(text)

	var word, sentence string
	is_word := false
	for _, c := range text {
		sentence += string(c)
		if IsAlpha(c) {
			word += string(c)
			is_word = true
			letters++
			if IsVowel(c) {
				vowels++
			} else {
				consonants++
			}
		} else if IsNumeric(c) {
			word += string(c)
			digits++
		} else if IsPunct(c) {
			if len(sentence) > 1 {
				sentences++
			}
			if is_word {
				words++
			}

			word = ""
			is_word = false
			sentence = ""

			specials++
		} else if IsSpace(c) {
			if is_word {
				words++
			}
			word = ""
			is_word = false
			spaces++
		} else {
			if is_word {
				words++
			}

			word = ""
			is_word = false

			specials++
		}
	}
	if is_word {
		words++
		sentences++
	}

	fmt.Printf("\nCharacters: %d\n", characters)
	fmt.Printf("Sentences: %d\n", sentences)
	fmt.Printf("Words: %d\n", words)
	fmt.Printf("Letters: %d\n", letters)
	fmt.Printf("Vowels: %d\n", vowels)
	fmt.Printf("Consonants: %d\n", consonants)
	fmt.Printf("Digits: %d\n", digits)
	fmt.Printf("Spaces: %d\n", spaces)
	fmt.Printf("Special Characters: %d\n", specials)
}
