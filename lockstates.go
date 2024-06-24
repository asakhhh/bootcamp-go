package bootcamp

type Lock struct {
	isLocked bool
}

func (l Lock) Lock() bool {
	if l.isLocked {
		return false
	}
	l.isLocked = true
	return true
}

func (l Lock) Unlock() bool {
	if !l.isLocked {
		return false
	}
	l.isLocked = false
	return true
}

func (l Lock) IsLocked() bool {
	return l.isLocked
}

func Newlock() Lock {
	return Lock{false}
}
