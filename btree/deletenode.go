package bootcamp

func (b *BTree) DeleteNode(value int) {
	if b == nil || b.Root == nil {
		return
	}
	curNode := b.Root
	for curNode != nil && curNode.Value != value {
		if value < curNode.Value {
			curNode = curNode.Left
		} else {
			curNode = curNode.Right
		}
	}
	if curNode == nil {
		return
	}

	if curNode.Left == nil && curNode.Right == nil {
		if curNode.Parent.Left == curNode {
			curNode.Parent.Left = nil
		} else {
			curNode.Parent.Right = nil
		}
		curNode.Parent = nil
	} else if curNode.Left != nil && curNode.Right == nil {
		if curNode.Parent.Left == curNode {
			curNode.Parent.Left = curNode.Left
		} else {
			curNode.Parent.Right = curNode.Left
		}
		curNode.Parent = nil
		curNode.Left = nil
	} else if curNode.Left == nil && curNode.Right != nil {
		if curNode.Parent.Left == curNode {
			curNode.Parent.Left = curNode.Right
		} else {
			curNode.Parent.Right = curNode.Right
		}
		curNode.Parent = nil
		curNode.Right = nil
	} else {
		newNode := curNode.Left
		for newNode.Right != nil {
			newNode = newNode.Right
		}
		curNode.Value = newNode.Value
		if newNode.Parent.Left == newNode {
			newNode.Parent.Left = newNode.Left
		} else {
			newNode.Parent.Right = newNode.Left
		}
	}
}
