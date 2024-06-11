package bootcamp

func Replace(s, old, new string) string {
	var t string
	for i := 0; i < len(s); {
		if i+len(old) <= len(s) && s[i:i+len(old)] == old {
			t += new
			i += len(old)
		} else {
			t += string(s[i])
			i++
		}
	}
	return t
}

// func main() {
// 	fmt.Println(Replace("Hello student!", "Hello", "Salem"))
// 	fmt.Println(Replace("banana", "a", "o"))
// }
