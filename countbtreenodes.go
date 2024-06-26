package bootcamp

import btree "bootcamp/btree"

func CountBtreeNodes(b *btree.BTree) int {
	return Size(b.Root)
}
