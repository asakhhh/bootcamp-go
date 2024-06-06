package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var a [20]int
	var mx *int = nil
	var mn *int = nil
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
		if mx == nil {
			mx = &a[i]
		} else if *mx < a[i] {
			mx = &a[i]
		}

		if mn == nil {
			mn = &a[i]
		} else if *mn > a[i] {
			mn = &a[i]
		}
	}

	fmt.Printf("%d %d\n", *mx, *mn)
}
