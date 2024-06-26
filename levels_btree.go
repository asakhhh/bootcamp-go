package bootcamp

import btree "bootcamp/btree"

func LevelsBTreeNode(n *btree.BTreeNode, lev int) int {
	if n.Left == nil && n.Right == nil {
		return lev
	}
	if n.Left != nil {
		return LevelsBTreeNode(n.Left, lev+1)
	}
	return LevelsBTreeNode(n.Right, lev+1)
}

func LevelsBtree(b *btree.BTree) int {
	if b == nil || b.Root == nil {
		return 0
	}
	return LevelsBTreeNode(b.Root, 1)
}
