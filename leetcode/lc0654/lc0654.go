package leetcode

import (
	"github.com/lyjgulu/algorithm/structures"
	"math"
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	minVal, index := math.MinInt64, 0
	for i, val := range nums {
		if val > minVal {
			minVal = val
			index = i
		}
	}

	root := &TreeNode{
		Val: nums[index],
	}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])

	return root
}
