package bootcamp

type Lock struct {
	isLocked *bool
}

func (l Lock) Lock() bool {
	if *l.isLocked {
		return false
	}
	*l.isLocked = true
	return true
}

func (l Lock) Unlock() bool {
	if !*l.isLocked {
		return false
	}
	*l.isLocked = false
	return true
}

func (l Lock) IsLocked() bool {
	return *l.isLocked
}

func NewLock() Lock {
	a := false
	return Lock{&a}
}

// func main() {
// 	lock := NewLock()
// 	fmt.Println(lock.IsLocked()) // false
// 	fmt.Println(lock.Unlock())   // false
// 	fmt.Println(lock.Lock())     // true
// 	fmt.Println(lock.Lock())     // false
// 	fmt.Println(lock.IsLocked()) // true
// 	fmt.Println(lock.Unlock())   // true
// }
