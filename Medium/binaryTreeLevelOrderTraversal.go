package main

import "fmt"

// Link: https://leetcode.com/problems/binary-tree-level-order-traversal/

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// if root is nil, return empty slice
	if root == nil {
		return [][]int{}
	}

	// create a result slice to store the result
	res := [][]int{}

	// create a queue to store the nodes
	queue := []*TreeNode{root}

	// create a level to store the level of the node
	level := 0

	// loop until the queue is empty
	for len(queue) > 0 {
		// add the value of the node to the result slice
		res = append(res, []int{})

		// loop until the queue is empty
		for _, node := range queue {
			// add the value of the current node to the result slice
			res[level] = append(res[level], node.Val)

			// slice the queue to remove the current node
			queue = queue[1:]

			// if the node has left child, add it to the temp queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// if the node has right child, add it to the temp queue
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// increment the level
		level++
	}

	// return the result
	return res
}

func main() {
	// test cases
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(levelOrder(root)) // [[3], [9, 20], [15, 7]]

	root = &TreeNode{
		Val: 1,
	}
	fmt.Println(levelOrder(root)) // [[1]]

	root = nil
	fmt.Println(levelOrder(root)) // []

}
