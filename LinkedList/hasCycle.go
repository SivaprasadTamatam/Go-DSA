package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// Helper function to create a linked list with a cycle
func createLinkedListWithCycle() *ListNode {
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Creating a cycle here

	return head
}

// Helper function to create a linked list without a cycle
func createLinkedListWithoutCycle() *ListNode {
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4

	return head
}

func main() {

	cycleList := createLinkedListWithCycle()
	fmt.Println("Test Case 1 (With Cycle):", hasCycle(cycleList)) // Expected Output: true

	noCycleList := createLinkedListWithoutCycle()
	fmt.Println("Test Case 2 (Without Cycle):", hasCycle(noCycleList)) // Expected Output: false
}
