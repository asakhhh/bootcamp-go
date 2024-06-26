package bootcamp

func (l *List) PushBefore(n *ListNode, v interface{}) {
	if n == nil {
		return
	}

	newNode := ListNode{nil, v}

	if l.Head == n {
		newNode.Next = l.Head
		l.Head = &newNode
		return
	}

	prevNode := l.Head
	for prevNode.Next != n {
		prevNode = prevNode.Next
	}
	prevNode.Next = &newNode
	newNode.Next = n
}
