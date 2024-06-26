package bootcamp

func (l *List) PushAfter(n *ListNode, v interface{}) {
	if n == nil {
		return
	}

	nextNode := n.Next
	newNode := ListNode{nil, v}
	n.Next = &newNode
	newNode.Next = nextNode
	if l.Tail == n {
		l.Tail = &newNode
	}
}
