package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head

	for curr != nil {
		nextNode := curr.Next
		curr.Next = pre
		pre = curr
		curr = nextNode
	}
	return pre
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, "	")
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	// Creating a sample linked list: 1 -> 2 -> 3 -> 4 -> 5 -> nil
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	fmt.Println("Original List:")
	printList(head)

	// Reverse the linked list
	head = reverseList(head)

	fmt.Println("Reversed List:")
	printList(head)

	// Making Original List
	head = reverseList(head)
	fmt.Println("Original List:")
	printList(head)

	head = reverseListRecursive(head)

	fmt.Println("Reversed List Using Recursive:")
	printList(head)
}
