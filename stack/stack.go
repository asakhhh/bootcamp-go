package bootcamp

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	Head, Tail *Node
}

type Stack struct {
	list *List
}

func NewStack() *Stack {
	return &Stack{&List{nil, nil}}
}

func (s *Stack) Push(v interface{}) {
	newNode := Node{v, nil}
	if s.list.Head == nil {
		s.list.Head = &newNode
		s.list.Tail = &newNode
	} else {
		newNode.Next = s.list.Head
		s.list.Head = &newNode
	}
}

func (s *Stack) Pop() interface{} {
	if s.list.Head == nil {
		return nil
	}
	curNode := s.list.Head
	s.list.Head = s.list.Head.Next
	if s.list.Head == nil {
		s.list.Tail = nil
	}
	curNode.Next = nil
	return curNode.Value
}
