package bootcamp

func IncrementIt(n *int) {
	if n != nil {
		(*n)++
	}
}

// func main() {
// 	num := 10

// 	IncrementIt(&num)
// 	fmt.Println(num)
// }
