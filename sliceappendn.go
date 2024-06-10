package bootcamp

// import "fmt"

func SliceAppendN(n int) []int {
	if n < 0 {
		return make([]int, 0)
	}

	var a []int

	for i := 0; i < n; i++ {
		a = append(a, i)
	}

	return a
}

// func main() {
// 	fmt.Println(SliceAppendN(5))
// }
