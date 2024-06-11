package bootcamp

func CountSubstr(s, substr string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if i+len(substr) <= len(s) && s[i:i+len(substr)] == substr {
			cnt++
		}
	}

	return cnt
}
