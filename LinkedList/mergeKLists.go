package main

import (
	"container/heap"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type NodeHeap []*ListNode

func (nh NodeHeap) Len() int           { return len(nh) }
func (nh NodeHeap) Less(i, j int) bool { return nh[i].Val < nh[j].Val }
func (nh NodeHeap) Swap(i, j int)      { nh[i], nh[j] = nh[j], nh[i] }

func (nh *NodeHeap) Push(x interface{}) {
	*nh = append(*nh, x.(*ListNode))
}

func (nh *NodeHeap) Pop() interface{} {
	old := *nh
	l := len(old)
	x := old[l-1]
	*nh = old[0 : l-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {

	h := &NodeHeap{}
	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for h.Len() > 0 {
		small := heap.Pop(h).(*ListNode)
		curr.Next = small
		curr = curr.Next
		if small.Next != nil {
			heap.Push(h, small.Next)
		}
	}
	return dummy.Next
}

// Helper: build linked list from slice
func BuildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head
}

// Helper: print list as slice
func ToSlice(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	// Example 1
	lists1 := []*ListNode{
		BuildList([]int{1, 4, 5}),
		BuildList([]int{1, 3, 4}),
		BuildList([]int{2, 6}),
	}
	merged1 := mergeKLists(lists1)
	fmt.Println(ToSlice(merged1)) // [1 1 2 3 4 4 5 6]

}
