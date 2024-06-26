package bootcamp

func (s *Stack) Peek() interface{} {
	if s.list.Head == nil {
		return nil
	}
	return s.list.Head.Value
}
