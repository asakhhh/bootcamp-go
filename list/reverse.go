package bootcamp

type ListNode struct {
	Next  *ListNode
	Value interface{}
}

type List struct {
	Head, Tail *ListNode
}

func NewList() *List {
	return &List{nil, nil}
}

func (l *List) Reverse() {
	curNode := l.Head
	var prevNode *ListNode

	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
	}
	a := l.Head
	l.Head = l.Tail
	l.Tail = a
}
