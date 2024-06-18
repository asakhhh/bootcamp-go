package bootcamp

func Contains(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}

func Pow10(num int) int {
	if num >= 1000 {
		return 3
	}
	if num >= 100 {
		return 2
	}
	if num >= 10 {
		return 1
	}
	return 0
}

func RomanToInt(s string) int {
	m := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
		"XX":   20,
		"XXX":  30,
		"XL":   40,
		"L":    50,
		"LX":   60,
		"LXX":  70,
		"LXXX": 80,
		"XC":   90,
		"C":    100,
		"CC":   200,
		"CCC":  300,
		"CD":   400,
		"D":    500,
		"DC":   600,
		"DCC":  700,
		"DCCC": 800,
		"CM":   900,
		"M":    1000,
		"MM":   2000,
		"MMM":  3000,
		"MMMM": 4000,
	}

	num := 0
	current_pos := 4
	i := 0
	for i < len(s) {
		found := false
		for l := 4; l > 0; l-- {
			if i+l <= len(s) && Contains(m, s[i:i+l]) {
				found = true
				add := m[s[i:i+l]]
				if current_pos <= Pow10(add) {
					return 0
				}
				current_pos = Pow10(add)
				// fmt.Printf("curpos == %d\n", current_pos)
				num += add
				i += l
				break
			}
		}
		if !found {
			return 0
		}
	}
	return num
}

// func main() {
// 	fmt.Println(RomanToInt("III"))     // 3
// 	fmt.Println(RomanToInt("IV"))      // 4
// 	fmt.Println(RomanToInt("IX"))      // 9
// 	fmt.Println(RomanToInt("LVIII"))   // 58
// 	fmt.Println(RomanToInt("MCMXCIV")) // 1994
// 	fmt.Println(RomanToInt(""))        // 0
// 	fmt.Println(RomanToInt("salem"))   // 0
// }
