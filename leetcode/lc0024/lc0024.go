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
// 递归法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	help := head.Next
	head.Next = swapPairs(help.Next)
	help.Next = head
	return help
}
