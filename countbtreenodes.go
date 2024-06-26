package bootcamp

import btree "bootcamp/btree"

func Size(n *btree.BTreeNode) int {
	if n == nil {
		return 0
	}
	return 1 + Size(n.Left) + Size(n.Right)
}

func CountBtreeNodes(b *btree.BTree) int {
	return Size(b.Root)
}
