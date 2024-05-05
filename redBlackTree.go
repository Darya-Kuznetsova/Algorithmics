package main

import "fmt"

// RED-BLACK tree is a BINARY tree:
// 1. A LEFT child is LESS than a parent
// 2. A RIGTH child is LARGER than a parent
// 3. Every node RED or BLACK
// 4. A ROOT is always BLACK
// 5. A NEW node is always RED
// 6. A RED note is always LEFT child
// 7. A RED node has only BLACK children
// 8. INSERTING and REMOVING are implemented with ROTATION or CHANGING COLOR
// 9. ROTATION can be LEFT or RED

func main() {
	newTree := NewRedBlackTree()

	newTree.insertRBTree(9)
	newTree.insertRBTree(7)
	newTree.insertRBTree(11)
	newTree.insertRBTree(5)
	newTree.insertRBTree(12)
	newTree.insertRBTree(8)
	newTree.insertRBTree(14)

	newTree.inorderRBTree()

}

type RBNode struct {
	color  string
	value  int
	left   *RBNode
	right  *RBNode
	parent *RBNode
}

type RBTree struct {
	root *RBNode
}

func NewRedBlackTree() *RBTree {
	return &RBTree{}
}

func (t *RBTree) insertRBTree(value int) {
	if t.root == nil {
		t.root = &RBNode{value: value, color: "BLACK"}
	} else {
		t.root.insertNode(value)
	}
}

func (n *RBNode) insertNode(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = &RBNode{value: value, color: "RED", parent: n}
			n.left.Fix()
		} else {
			n.left.insertNode(value)
		}
	} else if value > n.value {
		if n.right == nil {
			n.right = &RBNode{value: value, color: "RED", parent: n}
			n.right.Fix()
		} else {
			n.right.insertNode(value)
		}
	}
}
func (n *RBNode) Fix() {
	if n.parent == nil {
		n.color = "BLACK"
		return
	}

	if n.parent.color == "BLACK" {
		return
	}

	uncle := n.getUncle()
	grandparent := n.getGrandparent()

	if uncle != nil && uncle.color == "RED" {
		n.parent.color = "BLACK"
		uncle.color = "BLACK"
		grandparent.color = "RED"
		grandparent.Fix()
		return
	}

	if n == n.parent.right && n.parent == grandparent.left {
		grandparent.rotateLeft()
		n = n.left
	} else if n == n.parent.left && n.parent == grandparent.right {
		grandparent.rotateRight()
		n = n.right
	}

	n.parent.color = "BLACK"
	grandparent.color = "RED"

	if n == n.parent.left {
		grandparent.rotateRight()
	} else {
		grandparent.rotateLeft()
	}
}

func (n *RBNode) getUncle() *RBNode {
	if n.parent == nil || n.parent.parent == nil {
		return nil
	}

	grandparent := n.parent.parent
	if n.parent == grandparent.left {
		return grandparent.right
	}

	return grandparent.left
}

func (n *RBNode) getGrandparent() *RBNode {
	if n.parent != nil {
		return n.parent.parent
	}

	return nil
}

func (n *RBNode) rotateLeft() {
	child := n.right
	n.right = child.left

	if child.left != nil {
		child.left.parent = n
	}
	child.parent = n.parent

	if n.parent == nil {
		n.parent = child
	} else if n == n.parent.left {
		n.parent.left = child
	} else {
		n.parent.right = child
	}

	child.left = n
	n.parent = child

}

func (n *RBNode) rotateRight() {
	child := n.left
	n.left = child.right

	if child.right != nil {
		child.right.parent = n
	}
	child.parent = n.parent

	if n.parent == nil {
		n.parent = child
	} else if n == n.parent.right {
		n.parent.right = child
	} else {
		n.parent.left = child
	}

	child.right = n
	n.parent = child

}
func (t *RBTree) inorderRBTree() {
	if t.root != nil {
		t.root.inorderNode()
	}
	fmt.Println()
}
func (n *RBNode) inorderNode() {
	if n != nil {
		n.left.inorderNode()
		fmt.Printf("%d ", n.value)
		n.right.inorderNode()
	}
}
