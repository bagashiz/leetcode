package main

import "fmt"

// Link: https://leetcode.com/problems/binary-tree-right-side-view/

// TreeNode is a node of a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	// if root is nil, return empty slice
	if root == nil {
		return []int{}
	}

	// create a result slice to store the result
	res := []int{}

	// create a queue to store the nodes
	queue := []*TreeNode{root}

	// create a level to store the level of the node
	level := 0

	// create a map to store the level and the value of the node
	levelMap := make(map[int]int)

	// loop until the queue is empty
	for len(queue) > 0 {
		// create a temp queue to store the nodes of the next level
		tempQueue := []*TreeNode{}

		// loop until the queue is empty
		for _, node := range queue {
			// if the level of the node is not in the map, add it to the map
			if _, ok := levelMap[level]; !ok {
				levelMap[level] = node.Val
			}

			// if the node has right child, add it to the temp queue
			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}

			// if the node has left child, add it to the temp queue
			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}
		}

		// add the value of the node to the result slice
		res = append(res, levelMap[level])

		// update the level and the queue
		level++
		queue = tempQueue
	}

	return res
}

func main() {
	// test cases
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(rightSideView(root)) // [1, 3, 4]
}
