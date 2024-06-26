package bootcamp

func (n *BTreeNode) NodeClear() {
	if n != nil {
		n.Left.NodeClear()
		n.Right.NodeClear()
		n.Left = nil
		n.Right = nil
	}
}

func (b *BTree) Clear() {
	b.Root.NodeClear()
}
