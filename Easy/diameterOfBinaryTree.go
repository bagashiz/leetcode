package main

import (
	"fmt"
	"math"
)

// Link: https://leetcode.com/problems/diameter-of-binary-tree/

// binary tree struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	// create a variable to store the result
	result := 0

	// call helper function to get the height of the tree
	helper(root, &result)

	// return the result
	return result
}

func helper(root *TreeNode, result *int) int {
	// if root is nil, return 0
	if root == nil {
		return 0
	}

	// set L and R to be the height of left and right subtree
	L := helper(root.Left, result)
	R := helper(root.Right, result)

	// update result
	*result = int(math.Max(float64(*result), float64(L+R)))

	// return the height of the current node
	return int(math.Max(float64(L), float64(R))) + 1
}

func main() {
	// test cases
	fmt.Println(diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))                          // 3
	fmt.Println(diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}))                                                                                                       // 1
	fmt.Println(diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}})) // 4
}
