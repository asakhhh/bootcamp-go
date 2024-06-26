package bootcamp

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	Head, Tail *Node
}

func (l *List) PushBack(v interface{}) {
	newnode := Node{v, nil}
	if l.Tail == nil {
		l.Head = &newnode
		l.Tail = &newnode
	} else {
		l.Tail.Next = &newnode
		l.Tail = &newnode
	}
}

func (l *List) PopFront() interface{} {
	if l.Head == nil {
		return nil
	}
	curNode := l.Head
	l.Head = curNode.Next
	if l.Head == nil {
		l.Tail = nil
	}
	return curNode.Value
}

type Queue struct {
	list *List
}

func NewQueue() *Queue {
	return &Queue{&List{nil, nil}}
}

func (q *Queue) Push(v interface{}) {
	q.list.PushBack(v)
}

func (q *Queue) Pop() interface{} {
	return q.list.PopFront()
}
