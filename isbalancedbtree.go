package bootcamp

import btree "bootcamp/btree"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Size(n *btree.BTreeNode) int {
	if n == nil {
		return 0
	}
	return 1 + Size(n.Left) + Size(n.Right)
}

func IsBalancedBtreeNode(n *btree.BTreeNode) bool {
	if n != nil {
		return abs(Size(n.Left)-Size(n.Right)) <= 1
	}
	return true
}

func IsBalancedBtree(b *btree.BTree) bool {
	return IsBalancedBtreeNode(b.Root)
}
