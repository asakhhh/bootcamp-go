package bootcamp

func (l *List) Reverse() {
	curNode := l.Head
	var prevNode *ListNode

	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
	}
	a := l.Head
	l.Head = l.Tail
	l.Tail = a
}
