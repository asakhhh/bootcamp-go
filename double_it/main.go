package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d", 2*a[i])
		if i != n-1 {
			fmt.Printf(" ")
		}
	}
}
