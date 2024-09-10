package tree

import "fmt"

func (n *node) bstSearch(v int) *node {
	if n == nil {
		return nil
	}
	if v == n.data {
		return n
	}
	if v < n.data.(int) {
		return n.left.bstSearch(v)
	}
	if v > n.data.(int) {
		return n.right.bstSearch(v)
	}
	return n
}

func (n *node) bstInsert(v int) *node {
	if n.data == nil {
		*n = node{data: v}
	}
	if v < n.data.(int) {
		if n.left == nil {
			n.left = &node{data: v}
		} else {
			n.left.bstInsert(v)
		}
	}
	if v > n.data.(int) {
		if n.right == nil {
			n.right = &node{data: v}
		} else {
			n.right.bstInsert(v)
		}
	}
	return n
}

func BSTDemo() {
	root := newNode(1)
	root.bstInsert(5)
	root.bstInsert(2)
	root.bstInsert(9)
	root.bstInsert(10)
	root.bstInsert(8)
	root.bstInsert(3)
	root.bstInsert(4)
	root.bstInsert(6)
	root.bstInsert(7)
	root.inorderTraversal()

	fmt.Println(root.bstSearch(4))

}
