package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	A := head
	B := head
	for A != nil && A.Next != nil {
		A = A.Next.Next
		B = B.Next
	}
	return B
}
