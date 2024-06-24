package firststruct

func (u User) PasswordReliability() string {
	crit1 := len(u.password) >= 8

	crit2, crit3, crit4, crit5 := false, false, false, false
	for _, c := range u.password {
		if c >= '0' && c <= '9' {
			crit4 = true
		} else if c >= 'a' && c <= 'z' {
			crit3 = true
		} else if c >= 'A' && c <= 'Z' {
			crit2 = true
		} else {
			crit5 = true
		}
	}

	res := 0
	if crit1 {
		res++
	}
	if crit2 {
		res++
	}
	if crit3 {
		res++
	}
	if crit4 {
		res++
	}
	if crit5 {
		res++
	}

	if res == 5 {
		return "strong"
	} else if res >= 3 {
		return "medium"
	} else if res >= 1 {
		return "easy"
	}
	return "undefined"
}
