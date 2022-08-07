package main

import "fmt"

// Link: https://leetcode.com/problems/merge-two-sorted-lists/

// single linked list struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	// if list1 is empty, return list2
	if list1 == nil {
		return list2
	}

	// if list2 is empty, return list1
	if list2 == nil {
		return list1
	}

	// if list1.Val < list2.Val
	if list1.Val < list2.Val {
		// merge list1.Next and list2 then return list1
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else { // if list1.Val >= list2.Val
		// merge list1 and list2.Next then return list2
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func main() {
	// test cases
	fmt.Println(mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}})) // [1,1,2,3,4,4]
	fmt.Println(mergeTwoLists(nil, nil))                                                                                     // []
	fmt.Println(mergeTwoLists(nil, &ListNode{0, nil}))                                                                       // [0]
}
