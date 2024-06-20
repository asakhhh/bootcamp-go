package main

import (
	"fmt"
	"math/rand"

	"github.com/alem-platform/ap"
)

func PrintString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func main() {
	num := int(rand.Int31n(100) + 1)

	inp := -1

	for inp != num {
		PrintString("Guess number: ")
		_, err := fmt.Scanf("%d", &inp)
		if err != nil {
			PrintString("Enter a valid number.\n")
		} else {
			if num < inp {
				PrintString("Lower\n")
			} else if num > inp {
				PrintString("Higher\n")
			}
		}
	}
	PrintString("Match, you win!\n")
}
