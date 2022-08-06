package main

import "fmt"

// Link: https://leetcode.com/problems/palindrome-linked-list/

// single linked list struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// if head is nil, return true
	if head == nil {
		return true
	}

	// find the middle of the linked list
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second half of the linked list
	var prev *ListNode
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}

	// compare the first half and the second half
	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head, prev = head.Next, prev.Next
	}

	return true
}

func main() {
	// test cases
	fmt.Println(isPalindrome(&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}))               // true
	fmt.Println(isPalindrome(&ListNode{1, &ListNode{2, nil}}))                                           // false
	fmt.Println(isPalindrome(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}})) // true

}
