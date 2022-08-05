package main

import "fmt"

// Link: https://leetcode.com/problems/reverse-linked-list/

// single linked list struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// create a variable to store the previous node
	var prev *ListNode

	// loop through the linked list if head is not nil
	for head != nil {
		// create a variable to store the next node
		next := head.Next
		// set the next node to the previous node
		head.Next = prev
		// set the previous node to the current node
		prev = head
		// set the current node to the next node
		head = next
	}
	return prev
}

func main() {
	// test cases
	fmt.Println(reverseList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}})) // [5,4,3,2,1]
	fmt.Println(reverseList(&ListNode{1, &ListNode{2, nil}}))                                           // [2,1]
	fmt.Println(reverseList(&ListNode{1, nil}))                                                         // [1]
	fmt.Println(reverseList(nil))                                                                       // []
}
