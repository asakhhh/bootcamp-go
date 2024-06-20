package bootcamp

import "github.com/alem-platform/ap"

func PrintNum(n, length int) {
	p, curlen := 1, 1
	for p*10 <= n {
		p *= 10
		curlen++
	}

	for i := 0; i < length-curlen; i++ {
		ap.PutRune(' ')
	}

	for p > 0 {
		ap.PutRune(rune('0' + (n/p)%10))
		p /= 10
	}
}

func PrintSpiral(n int) {
	y := (n - 1) / 2
	x := y
	dy := []int{0, 1, 0, -1}
	dx := []int{1, 0, -1, 0}
	curdir := 0
	curmoved := 0

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}

	for i := 0; i < n*n; i++ {
		a[y][x] = i
		if curmoved < curdir/2+1 {
			curmoved++
		} else {
			curmoved = 1
			curdir++
		}
		x += dx[curdir%4]
		y += dy[curdir%4]
	}

	length, p := 1, 1
	for p*10 < n*n {
		p *= 10
		length++
	}

	for i, v := range a {
		for ii, vv := range v {
			PrintNum(vv, length)
			if ii != n-1 {
				ap.PutRune(' ')
			}
		}
		if i != n-1 {
			ap.PutRune('\n')
		}
	}
}

// func main() {
// 	PrintSpiral(4)
// }
