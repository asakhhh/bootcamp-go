package bootcamp

func (l *List) RemoveIf(fn func(n *ListNode) bool) {
	curNode := l.Head
	var prevNode *ListNode

	for curNode != nil {
		if fn(curNode) {
			nextNode := curNode.Next
			if prevNode != nil {
				prevNode.Next = nextNode
			}
			if l.Head == curNode {
				l.Head = l.Head.Next
			}
			if l.Tail == curNode {
				l.Tail = prevNode
			}
		} else {
			prevNode = curNode
		}
		curNode = curNode.Next
	}
}
