package bootcamp

func (n *BTreeNode) NodeClear() {
	if n != nil {
		n.Left.NodeClear()
		n.Right.NodeClear()
		n = nil
	}
}

func (b *BTree) Clear() {
	b.Root.NodeClear()
}
