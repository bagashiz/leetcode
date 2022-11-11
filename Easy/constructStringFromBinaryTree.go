package main

import (
	"fmt"
	"strconv"
)

// Link: https://leetcode.com/problems/construct-string-from-binary-tree/

// binary tree struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

  // if root does not have right branch
func tree2str(root *TreeNode) string {
	// if root is nil, return empty string
	if root == nil {
		return ""
	}

	// if root does not have right and left branch
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	}

  // if root does not have right branch
	if root.Right == nil {
    // print value of root and call tree2str() for left branch recursively
		return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")"
	}

  // if root does not have left branch
  // print value of root and call tree2str() for right branch and left recursively
	return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")(" + tree2str(root.Right) + ")"
}

func main() {
  // test cases
  fmt.Println(tree2str(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, nil}})) // "1(2(4))(3)"
  fmt.Println(tree2str(&TreeNode{1, &TreeNode{2, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}})) // "1(2()(4))(3)"
}
