package bootcamp

func (s *Stack) Len() int {
	curNode := s.list.Head
	ln := 0

	for curNode != nil {
		ln++
		curNode = curNode.Next
	}

	return ln
}
