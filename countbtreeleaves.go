package bootcamp

import btree "bootcamp/btree"

func LeafCount(n *btree.BTreeNode) int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return 1
	}
	return LeafCount(n.Left) + LeafCount(n.Right)
}

func CountBtreeLeaves(b *btree.BTree) int {
	return LeafCount(b.Root)
}
