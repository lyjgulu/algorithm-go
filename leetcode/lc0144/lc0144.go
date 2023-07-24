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

/*// 递归1
func preorderTraversal(root *TreeNode) []int {
    var res []int
    if root != nil {
        res = append(res, root.Val)
        tmp := preorderTraversal(root.Left)
        for _, t := range tmp {
            res = append(res, t)
        }
        tmp = preorderTraversal(root.Right)
        for _, t := range tmp {
            res = append(res, t)
        }
    }
    return res
}*/

/*// 递归2
func preorderTraversal(root *TreeNode) []int {
    var res []int
    preorder(root, &res)
    return res
}

func preorder(root *TreeNode, res *[]int) {
    for root != nil {
        *res = append(*res, root.Val)
        preorder(root.Left, res)
        preorder(root.Right, res)
    }
}*/

// 迭代
func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
