package main

import "fmt"

// Link: https://leetcode.com/problems/root-equals-sum-of-children/

// binary tree struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	// if root has two children and root equals sum of children, return true
	if root.Left.Val+root.Right.Val == root.Val {
		return true
	}
	return false
}

func main() {
	// test cases
	fmt.Println(checkTree(&TreeNode{Val: 10, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}})) // true
	fmt.Println(checkTree(&TreeNode{Val: 5, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}))  // false
}
