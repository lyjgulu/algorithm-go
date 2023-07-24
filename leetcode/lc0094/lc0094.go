package leetcode

import (
	"math"

	"github.com/lyjgulu/algorithm-go/structures"
)

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
func isValidBST(root *TreeNode) bool {
	var helper func(root *TreeNode, max int, min int) bool
	helper = func(root *TreeNode, max, min int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}
		return helper(root.Left, root.Val, min) && helper(root.Right, max, root.Val)
	}
	return helper(root, math.MaxInt64, math.MinInt64)
}
