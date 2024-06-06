package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var r rune
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &r)
		fmt.Printf("%d", r)
		if i != n-1 {
			fmt.Printf(" ")
		}
	}
}
