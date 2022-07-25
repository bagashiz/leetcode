package main

import "fmt"

// Link: https://leetcode.com/problems/symmetric-tree/

// binary tree struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// if root is nil, return true
	if root == nil {
		return true
	}

	// call helper function
	return isSymmetricHelper(root.Left, root.Right)
}

// isSymmetricHelper checks if the left and right subtrees are symmetric
func isSymmetricHelper(left, right *TreeNode) bool {
	// if left and right are nil, return true
	if left == nil && right == nil {
		return true
	}

	// if only one of left and right is nil, return false
	if left == nil || right == nil {
		return false
	}

	// if left and right have different values, return false
	if left.Val != right.Val {
		return false
	}

	// recurse on left and right subtrees
	return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
}

func main() {
	// test cases
	fmt.Println(isSymmetric(nil))                                                                                                                                                                        // true
	fmt.Println(isSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}})) // true
	fmt.Println(isSymmetric(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}))                                                   // false

}
