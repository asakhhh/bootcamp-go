package bootcamp

func MagicGrowthN(n int) []string {
	if n < 1 || n > 10 {
		return make([]string, 0)
	}
	if n == 1 {
		return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	}

	var res []string
	prev := MagicGrowthN(n - 1)
	for i := 0; i <= 9; i++ {
		ind := 0
		for ind < len(prev) && int(rune(prev[ind][0])-'0') <= i {
			ind++
		}
		for ind < len(prev) {
			t := string('0' + i)
			res = append(res, t+prev[ind])
			ind++
		}
	}

	return res
}

// func main() {
// 	fmt.Println(MagicGrowthN(1))  // ["1", "2", "3", "4", "5", "6", "7", "8", "9"]
// 	fmt.Println(MagicGrowthN(2))  // ["01", "02", ... "08", "09", "12", "13" ... "78", "79", "89"]
// 	fmt.Println(MagicGrowthN(3))  // ["012", "013", ... "089", "123", "123" ... "678", "679", "789"]
// 	fmt.Println(MagicGrowthN(9))  // ["012345678", "012345679", "012345689", ..., "123456789"]
// 	fmt.Println(MagicGrowthN(10)) // ["0123456789"]
// }
