package main

import "fmt"

// ListNode defines a node in the linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle checks if the linked list has a cycle.
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fmt.Println(fast)
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func main() {
	// Example usage.
	node4 := &ListNode{Val: -4}
	node3 := &ListNode{Val: 0, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 3, Next: node2}

	// Creating a cycle for the test: node4.Next = node2
	node4.Next = node2

	// Test the function
	fmt.Println("Does the linked list have a cycle? ", hasCycle(head))
}
