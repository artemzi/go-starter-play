// Package merge solve:
// https://leetcode.com/problems/merge-two-sorted-lists/description/
// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes of the first two lists.

// Example:
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
package merge

// ListNode is a definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (r *ListNode) {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		r = l1
		r.Next = mergeTwoLists(l1.Next, l2)
	} else {
		r = l2
		r.Next = mergeTwoLists(l2.Next, l1)
	}
	return
}
