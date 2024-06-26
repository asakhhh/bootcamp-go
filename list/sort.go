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

func (l *List) Len() int {
	ln := 0
	curNode := l.Head
	for curNode != nil {
		ln++
		curNode = curNode.Next
	}
	return ln
}

func (l *List) Clear() {
	curNode := l.Head

	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = nil
		curNode = nextNode
	}
	l.Head = nil
	l.Tail = nil
}

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

func (l *List) Sort(fn func(a *ListNode, b *ListNode) int) {
	if l.Len() <= 1 {
		return
	}
	n := l.Len()

	al, bl := NewList(), NewList()
	curNode := l.Head
	for i := 0; i < n/2; i++ {
		al.PushBack(curNode.Value)
		curNode = curNode.Next
	}
	for i := n / 2; i < n; i++ {
		bl.PushBack(curNode.Value)
		curNode = curNode.Next
	}
	al.Sort(fn)
	bl.Sort(fn)

	l.Clear()
	aNode, bNode := al.Head, bl.Head
	for aNode != nil && bNode != nil {
		if fn(aNode, bNode) < 0 {
			l.PushBack(aNode.Value)
			aNode = aNode.Next
		} else {
			l.PushBack(bNode.Value)
			bNode = bNode.Next
		}
	}
	for aNode != nil {
		l.PushBack(aNode.Value)
		aNode = aNode.Next
	}
	for bNode != nil {
		l.PushBack(bNode.Value)
		bNode = bNode.Next
	}
}
