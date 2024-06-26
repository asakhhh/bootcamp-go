package bootcamp

import tree "bootcamp/btree"

func GetMax(b *tree.BTree) *tree.BTreeNode {
	curNode := b.Root
	if curNode == nil {
		return nil
	}
	for curNode.Right != nil {
		curNode = curNode.Right
	}
	return curNode
}
