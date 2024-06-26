package bootcamp

func PostOrderTraversalNode(n *BTreeNode, fn func(n *BTreeNode)) {
	if n != nil {
		PreOrderTraversalNode(n.Left, fn)
		PreOrderTraversalNode(n.Right, fn)
		fn(n)
	}
}

func (b *BTree) PostOrderTraversal(fn func(n *BTreeNode)) {
	PostOrderTraversalNode(b.Root, fn)
}
