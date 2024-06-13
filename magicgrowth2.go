package bootcamp

func MagicGrowth2() []string {
	var res []string
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			var t string
			t += string('0' + i)
			t += string('0' + j)
			res = append(res, t)
		}
	}
	return res
}
