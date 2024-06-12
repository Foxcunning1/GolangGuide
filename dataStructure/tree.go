package main

import "fmt"

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Insert inserts a new value into the binary tree
func (n *TreeNode) Insert(value int) {
	if n == nil {
		return
	}
	if value < n.Value {
		if n.Left == nil {
			n.Left = &TreeNode{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search searches for a value in the binary tree
func (n *TreeNode) Search(value int) *TreeNode {
	if n == nil || n.Value == value {
		return n
	}
	if value < n.Value {
		return n.Left.Search(value)
	}
	return n.Right.Search(value)
}

// Preorder traverses the tree in preorder
func (n *TreeNode) Preorder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Value)
	n.Left.Preorder()
	n.Right.Preorder()
}

// Inorder traverses the tree in inorder
func (n *TreeNode) Inorder() {
	if n == nil {
		return
	}
	n.Left.Inorder()
	fmt.Printf("%d ", n.Value)
	n.Right.Inorder()
}

// Postorder traverses the tree in postorder
func (n *TreeNode) Postorder() {
	if n == nil {
		return
	}
	n.Left.Postorder()
	n.Right.Postorder()
	fmt.Printf("%d ", n.Value)
}

func main() {
	root := &TreeNode{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(2)
	root.Insert(7)
	root.Insert(12)
	root.Insert(20)

	fmt.Println("Preorder traversal:")
	root.Preorder()
	fmt.Println("")

	fmt.Println("Inorder traversal:")
	root.Inorder()
	fmt.Println()

	fmt.Println("Postorder traversal:")
	root.Postorder()
	fmt.Println()

	searchValue := 7
	if foundNode := root.Search(searchValue); foundNode != nil {
		fmt.Printf("Found node with value: %d\n", foundNode.Value)
	} else {
		fmt.Printf("Node with value %d not found\n", searchValue)
	}
}
