package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	s := make([]rune, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &s[i])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d", int(0+s[i]))
		if i != n-1 {
			fmt.Printf(" ")
		}
	}
}
