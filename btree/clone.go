package bootcamp

func CloneNode(src, dst *BTreeNode) {
	dst.Value = src.Value
	if src.Left != nil {
		dst.Left = &BTreeNode{dst, nil, nil, 0}
		CloneNode(src.Left, dst.Left)
	}
	if src.Right != nil {
		dst.Right = &BTreeNode{dst, nil, nil, 0}
		CloneNode(src.Right, dst.Right)
	}
}

func (b *BTree) Clone() *BTree {
	if b == nil {
		return nil
	}

	newTree := NewBTree()
	if b.Root != nil {
		newTree.Root = &BTreeNode{nil, nil, nil, 0}
		CloneNode(b.Root, newTree.Root)
	}
	return newTree
}
