package main

import (
	"fmt"
	"math"
)

// Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/

// Binary Tree Struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// if the binary tree is empty, return 0
	if root == nil {
		return 0
	}

	// if the binary tree is not empty, return the most deep level of the binary tree node + 1 using the max function
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))) + 1)
}

func main() {
	// test cases
	fmt.Println(maxDepth(nil))                                                                                                                       // 0
	fmt.Println(maxDepth(&TreeNode{Val: 3, Right: &TreeNode{Val: 20, Right: &TreeNode{Val: 7}, Left: &TreeNode{Val: 15}}, Left: &TreeNode{Val: 9}})) // 3
	fmt.Println(maxDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}))                                                                                // 2
}
