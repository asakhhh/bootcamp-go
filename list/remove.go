package bootcamp

func (l *List) Remove(n *ListNode) {
	if l.Head == n { // len == 1
		l.Head = l.Head.Next
		if l.Tail == n {
			l.Tail = nil
		}
		return
	}
	// len > 1

	prevNode := l.Head
	for prevNode != nil && prevNode.Next != n {
		prevNode = prevNode.Next
	}
	if prevNode == nil {
		return
	}
	if l.Tail == n {
		l.Tail = prevNode
	}

	nextNode := n.Next
	prevNode.Next = nextNode
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
