package bootcamp

func (l *List) Len() int {
	ln := 0
	curNode := l.Head
	for curNode != nil {
		ln++
		curNode = curNode.Next
	}
	return ln
}
