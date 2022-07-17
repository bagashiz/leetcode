package main

import "fmt"

// Link: https://leetcode.com/problems/binary-tree-inorder-traversal/

// create TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// if root is nil, return empty array
	if root == nil {
		return []int{}
	}

	// create empty array to store inorder traversal
	var inorder []int

	// create stack to store nodes
	var stack []*TreeNode

	// create pointer to current node
	current := root

	// loop until current is nil
	for current != nil {
		// push current node to stack
		stack = append(stack, current)

		// set current to current's left child
		current = current.Left
	}

	// loop until stack is empty
	for len(stack) > 0 {
		// pop node from stack
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// append current node's value to inorder traversal
		inorder = append(inorder, current.Val)

		// set current to current's right child
		current = current.Right

		// loop until current is nil
		for current != nil {
			// push current node to stack
			stack = append(stack, current)

			// set current to current's left child
			current = current.Left
		}
	}

	return inorder
}

func main() {
	// test cases
	fmt.Println(inorderTraversal(nil))                                                          // []
	fmt.Println(inorderTraversal(&TreeNode{Val: 1}))                                            // [1]
	fmt.Println(inorderTraversal(&TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}})) // [1, 3, 2]
}
