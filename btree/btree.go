package bootcamp

type BTreeNode struct {
	Parent      *BTreeNode
	Left, Right *BTreeNode
	Value       int
}

type BTree struct {
	Root *BTreeNode
}

func NewBTree() *BTree {
	return &BTree{&BTreeNode{nil, nil, nil, 0}}
}
