package bootcamp

func ArraySum(arr [20]int) int {
	sm := 0

	for i := 0; i < 20; i++ {
		sm += arr[i]
	}

	return sm
}
