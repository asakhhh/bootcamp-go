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
	return &BTree{nil}
}

func (b *BTree) ReplaceOrInsert(v int) {
	curNode := b.Root
	if curNode == nil {
		b.Root = &BTreeNode{nil, nil, nil, v}
		return
	}
	for {
		if v == curNode.Value {
			return
		}
		if v < curNode.Value {
			if curNode.Left == nil {
				newNode := BTreeNode{curNode, nil, nil, v}
				curNode.Left = &newNode
				return
			}
			curNode = curNode.Left
		} else {
			if curNode.Right == nil {
				newNode := BTreeNode{curNode, nil, nil, v}
				curNode.Right = &newNode
				return
			}
			curNode = curNode.Right
		}
	}
}
