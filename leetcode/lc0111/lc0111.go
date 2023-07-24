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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := math.MaxInt32
	if root.Left != nil {
		min = compareMin(min, minDepth(root.Left))
	}
	if root.Right != nil {
		min = compareMin(min, minDepth(root.Right))
	}
	return min + 1
}

func compareMin(paraOne int, paraTwo int) int {
	if paraOne > paraTwo {
		return paraTwo
	}
	return paraOne
}
