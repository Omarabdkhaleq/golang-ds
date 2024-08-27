package tree

import (
	"fmt"
	"math"
)

type node struct {
	data  interface{}
	left  *node
	right *node
}

func newNode(v interface{}) *node {
	return &node{
		data:  v,
		left:  nil,
		right: nil,
	}
}

func (n *node) insertLeft(target *node) {
	if n.left != nil {
		panic("left node already exist")
	}
	n.left = target
}

func (n *node) insertRight(target *node) {
	if n.right != nil {
		panic("right node already exist")
	}
	n.right = target
}

func (n *node) depth() int {
	if n == nil {
		return 0
	}
	return int(math.Max(float64(n.left.depth()), float64(n.right.depth())) + 1)
}

func (n *node) nodesCount() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.nodesCount() + n.right.nodesCount()
}

func (n *node) orderTraversal() {
	if n == nil {
		return
	}
	fmt.Println("node", n.data)
	n.left.orderTraversal()
	n.right.orderTraversal()
}

func (n *node) inorderTraversal() {
	if n == nil {
		return
	}
	n.left.inorderTraversal()
	fmt.Println("node", n.data)
	n.right.inorderTraversal()
}

func (n *node) postorderTraversal() {
	if n == nil {
		return
	}
	n.left.postorderTraversal()
	n.right.postorderTraversal()
	fmt.Println("node", n.data)
}

func (n *node) isFullBinaryTree() bool {
	if n == nil {
		return true
	}

	if n.left == nil && n.right == nil {
		return true
	}

	if n.left != nil && n.right != nil {
		return n.left.isFullBinaryTree() && n.right.isFullBinaryTree()
	}

	return false
}

func (n *node) isComplete(index int, nodesCount int) bool {
	if n == nil {
		return true
	}

	if index >= nodesCount {
		return false
	}

	return n.left.isComplete(2*index+1, nodesCount) && n.right.isComplete(2*index+2, nodesCount)
}

func (n *node) isPerfect() bool {
	depth := n.depth()
	lvl := 0
	if n == nil {
		return true
	}

	if n.left == nil && n.right == nil {
		return depth == lvl+1
	}

	if n.left == nil || n.right == nil {
		return false
	}

	return n.left.isPerfect() && n.right.isPerfect()
}

func (n *node) isBalanced(h *int) bool {
	var lh, rh int = 0, 0

	if n == nil {
		*h = 0
		return true
	}

	l := n.left.isBalanced(&lh)
	r := n.right.isBalanced(&rh)

	*h = int(math.Max(float64(lh), float64(rh))) + 1

	if math.Abs(float64(lh-rh)) >= 2 {
		return false
	}

	return l && r
}

func BinaryTreeDemo() {
	root := newNode(1)
	l := newNode(2)
	r := newNode(3)
	ll := newNode(4)
	lr := newNode(5)
	rl := newNode(6)
	rr := newNode(7)

	root.insertLeft(l)
	root.insertRight(r)
	l.insertLeft(ll)
	l.insertRight(lr)
	r.insertLeft(rl)
	r.insertRight(rr)

	fmt.Println("Order Traversal")
	root.orderTraversal()
	fmt.Println("Inorder Traversal")
	root.inorderTraversal()
	fmt.Println("Postorder Traversal")
	root.postorderTraversal()

	fmt.Println("Nodes Count:", root.nodesCount())
	fmt.Println("Depth:", root.depth())

	fmt.Println("Full Binary:", root.isFullBinaryTree())
	fmt.Println("Is Perfect:", root.isPerfect())
	fmt.Println("Is Complete:", root.isComplete(0, root.nodesCount()))

	var x int = 0
	fmt.Println("Is Balanced:", root.isBalanced(&x))
}
