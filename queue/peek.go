package bootcamp

func (q *Queue) Peek() interface{} {
	if q.list.Head == nil {
		return nil
	}
	return q.list.Head.Value
}
