package leetcode

import (
	"github.com/lyjgulu/algorithm/structures"
)

type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) (int, int) {
	if head == nil || head.Next == nil {
		return 0, 0
	}
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return 0, 0
		}
		fast = fast.Next.Next
		if slow == fast {
			help := head
			for help != slow {
				help = help.Next
				slow = slow.Next
			}
			return 1, help.Val
		}
	}
	return 0, 0
}
