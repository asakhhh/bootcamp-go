package bootcamp

import btree "bootcamp/btree"

func DeleteBtreeLeavesNode(n *btree.BTreeNode) {
	if n == nil || (n.Left == nil && n.Right == nil) {
		return
	}
	if n.Left != nil {
		leftNode := n.Left
		if leftNode.Left == nil && leftNode.Right == nil {
			leftNode.Parent = nil
			n.Left = nil
		} else {
			DeleteBtreeLeavesNode(leftNode)
		}
	}
	if n.Right != nil {
		rightNode := n.Right
		if rightNode.Left == nil && rightNode.Right == nil {
			rightNode.Parent = nil
			n.Right = nil
		} else {
			DeleteBtreeLeavesNode(rightNode)
		}
	}
}

func DeleteBtreeLeaves(b *btree.BTree) {
	if b == nil || b.Root == nil {
		return
	}
	if b.Root.Left == nil && b.Root.Right == nil {
		b.Root = nil
	} else {
		DeleteBtreeLeavesNode(b.Root)
	}
}
