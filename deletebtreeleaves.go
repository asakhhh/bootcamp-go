package bootcamp

import btree "bootcamp/btree"

func DeleteBtreeLeavesNode(n *btree.BTreeNode) {
	if n != nil {
		n.Parent = nil
		DeleteBtreeLeavesNode(n.Left)
		DeleteBtreeLeavesNode(n.Right)
		n.Left = nil
		n.Right = nil
	}
}

func DeleteBtreeLeaves(b *btree.BTree) {
	DeleteBtreeLeavesNode(b.Root)
}
