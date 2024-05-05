package main

import "fmt"

// A BINARY TREE - every node has no more two children.
// A left child is less than a parent. A right child is larger than a parent.

func main() {
	root := Node{value: 9}

	root.left_child = &Node{value: 4}
	root.right_child = &Node{value: 12}

	root.insert(3)
	root.insert(2)
	root.insert(11)
	root.insert(10)
	root.insert(13)

	root.preorder()

	// fmt.Println("Depth: ", root.depth())
	// fmt.Println("Number of nodes: ", root.nodes_number())
}

type Node struct {
	value       int
	left_child  *Node
	right_child *Node
}

// INSERT (A ROOT ALREADY EXISTS):

func (n *Node) insert(number int) {
	if number < n.value {
		if n.left_child == nil {
			n.left_child = &Node{value: number, left_child: nil, right_child: nil}
		} else {
			n.left_child.insert(number)
		}
	} else {
		if n.right_child == nil {
			n.right_child = &Node{value: number, left_child: nil, right_child: nil}
		} else {
			n.right_child.insert(number)
		}
	}
}

// PREORDER TRAVERSAL (root -> left child -> right child):

func (n *Node) preorder() {
	if n != nil {
		fmt.Print(n.value, " ")
		n.left_child.preorder()
		n.right_child.preorder()
	}
}

// A DEPTH:

func (n *Node) depth() int {
	count := 0
	if n == nil {
		return count
	}
	if n.left_child.depth() > n.right_child.depth() {
		count = n.left_child.depth()
	} else {
		count = n.right_child.depth()
	}
	return count + 1
}

// NUMBER OF NODES:

func (n *Node) nodes_number() int {
	if n == nil {
		return 0
	} else {
		return n.left_child.nodes_number() + n.right_child.nodes_number() + 1
	}
}
