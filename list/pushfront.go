package bootcamp

func (l *List) PushFront(v interface{}) {
	newnode := ListNode{nil, v}
	if l.Head == nil {
		l.Head = &newnode
		l.Tail = &newnode
	} else {
		newnode.Next = l.Head
		l.Head = newnode.Next
	}
}
