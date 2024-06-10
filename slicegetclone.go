package bootcamp

// import "fmt"

func SliceGetClone(src []int) []int {
	a := make([]int, cap(src))

	for i := 0; i < len(src); i++ {
		a[i] = src[i]
	}

	return a
}
