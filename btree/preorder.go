package bootcamp

func PreOrderTraversalNode(n *BTreeNode, fn func(n *BTreeNode)) {
	if n != nil {
		fn(n)
		PreOrderTraversalNode(n.Left, fn)
		PreOrderTraversalNode(n.Right, fn)
	}
}

func (b *BTree) PreOrderTraversal(fn func(n *BTreeNode)) {
	PreOrderTraversalNode(b.Root, fn)
}
