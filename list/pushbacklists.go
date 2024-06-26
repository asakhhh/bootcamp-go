package bootcamp

func (l *List) PushBackLists(lists ...*List) {
	for _, v := range lists {
		if l.Head == nil {
			l.Head = v.Head
			l.Tail = v.Tail
		} else {
			l.Tail.Next = v.Head
			l.Tail = v.Tail
		}
	}
}
