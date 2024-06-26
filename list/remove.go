package bootcamp

func (l *List) Remove(n *ListNode) {
	if l.Head == n {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		return
	}

	prevNode := l.Head
	for prevNode != nil && prevNode != n {
		prevNode = prevNode.Next
	}
	if prevNode == nil {
		return
	}

	nextNode := n.Next
	prevNode.Next = nextNode
	if l.Head == n {
		l.Head = n.Next
	}
	if l.Tail == n {
		l.Tail = prevNode
	}
}

// func main() {
// 	l := NewList()
// 	l.PushBack(10)
// 	l.PushBack(20)
// 	l.PushBack(30)

// 	node := l.Head.Next // Node with value 20
// 	l.Remove(node)

// 	node = l.Head
// 	for node != nil {
// 		fmt.Println(node.Value)
// 		node = node.Next
// 	}
// 	// Output:
// 	// 10
// 	// 30
// }
