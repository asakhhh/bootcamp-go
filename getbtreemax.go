package bootcamp

func GetMax(b *tree.BTree) *BTreeNode {
	curNode := b.Root
	if curNode == nil {
		return nil
	}
	for curNode.Right != nil {
		curNode = curNode.Right
	}
	return curNode
}
