package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Scanf("%d %d", &num1, &num2)

	fmt.Printf("%d %d %d ", num1+num2, num1-num2, num1*num2)
	if num2 == 0 {
		fmt.Printf("F\n")
	} else {
		fmt.Printf("%d\n", num1/num2)
	}
}
