package template

import "github.com/lyjgulu/algorithm/structures"

type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listNodeTraverse1(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		// 迭代访问 p.val
	}
}

func listNodeTraverse2(head *ListNode) {
	// 递归访问 head.val
	listNodeTraverse2(head.Next)
}
