package main

import (
	"fmt"
)

// ListNode is the structure for list nodes
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates removes duplicates from a sorted linked list
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

// Helper function to create a linked list from a slice
func createLinkedList(elements []int) *ListNode {
	if len(elements) == 0 {
		return nil
	}

	head := &ListNode{Val: elements[0]}
	current := head
	for _, val := range elements[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example input
	input := []int{1, 1, 2}

	// Create a linked list from the input
	head := createLinkedList(input)

	// Remove duplicates
	head = deleteDuplicates(head)

	// Print the result
	printLinkedList(head)
}
