package main

import (
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Scanf("%d %d", &num1, &num2)

	if num2 == 0 {
		fmt.Printf("%d %d %d F", num1+num2, num1-num2, num1*num2)
	} else {
		fmt.Printf("%d %d %d %d", num1+num2, num1-num2, num1*num2, num1/num2)
	}
}
