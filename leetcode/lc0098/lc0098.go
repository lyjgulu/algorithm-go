package leetcode

import "github.com/lyjgulu/algorithm-go/structures"

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归
/*
func inorderTraversal(root *TreeNode) []int {
    var res []int
    inorder(root, &res)
    return res
}

func inorder(node *TreeNode, output *[]int) {
    if node != nil {
        inorder(node.Left, output)
        *output = append(*output, node.Val)
        inorder(node.Right, output)
    }
}
*/

// 迭代(自定义栈)
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
