package template

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func MultipleTreeTraverse(node *Node) {
	for _, child := range node.Children {
		MultipleTreeTraverse(child)
	}
}
