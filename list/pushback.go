package bootcamp

func (l *List) PushBack(v interface{}) {
	newnode := ListNode{v, nil}
	if l.Head == nil {
		l.Head = &newnode
		l.Tail = &newnode
	} else {
		l.Tail.Next = &newnode
		l.Tail = &newnode
	}
}
