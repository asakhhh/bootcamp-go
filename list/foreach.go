package bootcamp

func (l *List) ForEach(fn func(n *ListNode)) {
	curNode := l.Head

	for curNode != nil {
		fn(curNode)
		curNode = curNode.Next
	}
}
