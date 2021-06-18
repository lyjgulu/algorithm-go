package leetcode

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	dfs(root, &res)
	return res
}

func dfs(root *Node, output *[]int) {
	if root == nil {
		return
	}
	for _, val := range root.Children {
		dfs(val, output)
	}
	*output = append(*output, root.Val)
}
