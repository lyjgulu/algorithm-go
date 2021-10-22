package template

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func MultipleTreeTraverse(node *Node) {
	for _, child := range node.Children {
		// 前序遍历需要的操作
		MultipleTreeTraverse(child)
		// 后序遍历需要的操作
	}
}
