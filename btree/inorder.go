package bootcamp

func InOrderTraversalNode(n *BTreeNode, fn func(n *BTreeNode)) {
	if n != nil {
		InOrderTraversalNode(n.Left, fn)
		fn(n)
		InOrderTraversalNode(n.Right, fn)
	}
}

func (b *BTree) InOrderTraversal(fn func(n *BTreeNode)) {
	InOrderTraversalNode(b.Root, fn)
}
