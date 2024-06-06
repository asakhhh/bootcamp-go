package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	s := make([]rune, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &s[i])
		if s[i] == '\n' {
			break
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d", s[i])
		if i != n-1 {
			fmt.Printf(" ")
		}
	}
}
