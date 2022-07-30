package main

import "fmt"

// Link: https://leetcode.com/problems/add-two-numbers/

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	// create variable to carry over the sum if it is greater than 10
	var carry int

	// create a dummy node to store the sum of the two lists
	dummy := &ListNode{}

	// create a pointer to the current node
	current := dummy

	// loop through the two lists until one of them is empty
	for l1 != nil || l2 != nil {
		// create a new node to store the carry over
		sum := carry

		// if l1 is not empty, add the value to the sum and move to the next node
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// if l2 is not empty, add the value to the sum and move to the next node
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// set the carry over to the remainder of the sum divided by 10
		carry = sum / 10

		// set the next node in current to sum modulo 10
		current.Next = &ListNode{Val: sum % 10}

		// set the current node to the next node
		current = current.Next
	}

	// if the carry is not 0, set the next node to the carry
	if carry != 0 {
		current.Next = &ListNode{Val: carry}
	}

	// return the dummy node
	return dummy.Next
}

func main() {
	// test cases
	fmt.Println(addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}))                                                                                                                              // 7->0->8
	fmt.Println(addTwoNumbers(&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}, &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}})) // 8->9->9->9->0->0->0->1
}
