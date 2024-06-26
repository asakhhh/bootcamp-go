package bootcamp

func (q *Queue) Peek() interface{} {
	return q.list.Head
}
