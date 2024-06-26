package bootcamp

func (q *Queue) Len() int {
	ln := 0
	curNode := q.list.Head
	for curNode != nil {
		ln++
		curNode = curNode.Next
	}
	return ln
}
