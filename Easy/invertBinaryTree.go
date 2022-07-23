package main

import "fmt"

// Link: https://leetcode.com/problems/invert-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	// check if root is nil
	if root == nil {
		return nil
	}
	// invert left and right nodes
	root.Left, root.Right = root.Right, root.Left

	// invert left and right nodes of left and right nodes by using recursion
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func main() {
	// test cases
	fmt.Println(invertTree(nil)) // nil
	fmt.Println(invertTree(
		&TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		})) // 4,7,2,9,6,3,1
	fmt.Println(invertTree(
		&TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		})) // 2,3,1
}
