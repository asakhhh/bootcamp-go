package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func main() {
	var num1, num2 int

	fmt.Scanf("%d %d", &num1, &num2)

	if num2 == 0 {
		fmt.Printf("%d %d %d F\n", num1+num2, num1-num2, num1*num2)
	} else {
		fmt.Printf("%d %d %d %d\n", num1+num2, num1-num2, num1*num2, num1/num2)
	}
	ap.PutRune('\n')
}
