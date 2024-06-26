package bootcamp

import tree "bootcamp/btree"

func GetMin(b *tree.BTree) *tree.BTreeNode {
	curNode := b.Root
	if curNode == nil {
		return nil
	}
	for curNode.Left != nil {
		curNode = curNode.Left
	}
	return curNode
}
