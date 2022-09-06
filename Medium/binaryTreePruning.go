package main

import "fmt"

// Link: https://leetcode.com/problems/binary-tree-pruning/

// binary tree struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	// if root is nil, return nil
	if root == nil {
		return nil
	}

	// call pruneTree recursively on left and right nodes
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	// if root is a leaf node and its value is 0
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		// return nil
		return nil
	}

	// return root
	return root
}

func main() {
	// test cases
	fmt.Println(pruneTree(&TreeNode{1, nil, &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}}}))                                                          // 1, nil, 0, nil, 1
	fmt.Println(pruneTree(&TreeNode{1, &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}}, &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}}})) // 1, nil, 1, nil, 1
}
