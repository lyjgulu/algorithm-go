package template

import (
	"github.com/lyjgulu/algorithm/structures"
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
func TreeTraverse(root *TreeNode) {
	// 前序遍历
	TreeTraverse(root.Left)
	// 中序遍历
	TreeTraverse(root.Right)
	// 后序遍历
}
