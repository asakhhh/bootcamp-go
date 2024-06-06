package bootcamp

import "fmt"

func ArrayPrint(arr [20]int) {
	for i := 0; i < 20; i++ {
		fmt.Printf("%d", arr[i])
		if i != 19 {
			fmt.Printf(" ")
		}
	}
}
