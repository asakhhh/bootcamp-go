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

// func main() {
// 	l := NewList()
// 	fmt.Println(l)
// 	// Output: &{<nil> <nil>}
// }
