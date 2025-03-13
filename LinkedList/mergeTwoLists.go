package main

import (
	"fmt"
)

// ListNode represents a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two sorted linked lists and returns the head of the merged list.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}

	curr := res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return res.Next
}

// Print linked list
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " -> ")
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {
	// Create example lists
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	l2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}

	// Merge the lists
	mergedList := mergeTwoLists(l1, l2)

	// Print the merged list
	fmt.Print("Merged List: ")
	printList(mergedList)
}
