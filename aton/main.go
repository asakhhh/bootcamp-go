package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var s string
	fmt.Scanf("%s", &s)

	for i := 0; i < n; i++ {
		fmt.Printf("%d", s[i])
		if i != n-1 {
			fmt.Printf(" ")
		}
	}
}
