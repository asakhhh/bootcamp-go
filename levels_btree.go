package bootcamp

import btree "bootcamp/btree"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func LevelsBTreeNode(n *btree.BTreeNode, lev int) int {
	if n.Left == nil && n.Right == nil {
		return lev
	}
	if n.Left != nil {
		lev = max(lev, LevelsBTreeNode(n.Left, lev+1))
	}
	return max(lev, LevelsBTreeNode(n.Right, lev+1))
}

func LevelsBtree(b *btree.BTree) int {
	if b == nil || b.Root == nil {
		return 0
	}
	return LevelsBTreeNode(b.Root, 1)
}
