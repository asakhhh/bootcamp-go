package bootcamp

func ApplyByLevelNode(n *BTreeNode, lev int, fn func(node *BTreeNode, level int)) {
	fn(n, lev)
	if n.Left != nil {
		ApplyByLevelNode(n.Left, lev+1, fn)
	}
	if n.Right != nil {
		ApplyByLevelNode(n.Right, lev+1, fn)
	}
}

func (b *BTree) ApplyByLevel(fn func(node *BTreeNode, level int)) {
	if b == nil || b.Root == nil {
		return
	}
	ApplyByLevelNode(b.Root, 0, fn)
}
