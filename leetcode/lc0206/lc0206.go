package leetcode

import "github.com/lyjgulu/algorithm-go/structures"

type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var help *ListNode
	for head != nil {
		next := head.Next
		head.Next = help
		help = head
		head = next
	}
	return help
}
