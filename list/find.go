package bootcamp

func (l *List) Find(fn func(v interface{}) bool) *ListNode {
	curNode := l.Head
	for curNode != nil && !fn(curNode.Value) {
		curNode = curNode.Next
	}
	return curNode
}
