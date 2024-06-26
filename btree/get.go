package bootcamp

func (b *BTree) Get(value int) *BTreeNode {
	curNode := b.Root
	for curNode != nil {
		if value == curNode.Value {
			return curNode
		} else if value < curNode.Value {
			curNode = curNode.Left
		} else {
			curNode = curNode.Right
		}
	}
	return nil
}
