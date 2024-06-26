package bootcamp

func (l *List) PushBack(v interface{}) {
	newnode := ListNode{nil, v}
	if l.Head == nil {
		l.Head = &newnode
		l.Tail = &newnode
	} else {
		l.Tail.Next = &newnode
		l.Tail = &newnode
	}
}

// func main() {
// 	l := NewList()
// 	l.PushBack(10)
// 	l.PushBack(20)
// 	l.PushBack(30)

// 	node := l.Head
// 	for node != nil {
// 		fmt.Println(node.Value)
// 		node = node.Next
// 	}
// 	// Output:
// 	// 10
// 	// 20
// 	// 30
// }
