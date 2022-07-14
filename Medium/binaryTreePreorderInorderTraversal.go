package main

import "fmt"

// Link: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

// Create a struct to represent a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder, inorder []int) *TreeNode {
	// if the length of the preorder is 0, return nil
	if len(preorder) == 0 {
		return nil
	}

	// get the root value from the preorder
	rootVal := &TreeNode{Val: preorder[0]}

	// loop through the inorder array
	for i := 0; i < len(inorder); i++ {
		// if the current value in the inorder array is equal to the root value,
		if inorder[i] == rootVal.Val {
			// get the left and right subtrees from the preorder array with recursive calls
			rootVal.Left = buildTree(preorder[1:i+1], inorder[:i])
			rootVal.Right = buildTree(preorder[i+1:], inorder[i+1:])

			// break out of the loop
			break
		}
	}

	return rootVal
}

func main() {
	// test cases
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})) // [3,9,20,null,null,15,7]
	fmt.Println(buildTree([]int{1, 2, 3}, []int{2, 1, 3}))                 // [1,2,3]
	fmt.Println(buildTree([]int{-1}, []int{-1}))                           // [-1]

}
