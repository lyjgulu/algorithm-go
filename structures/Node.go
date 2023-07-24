package structures

import "github.com/lyjgulu/algorithm-go/structures"

type Node struct {
	Val      int
	Children []*Node
}

// Int2NaryNode 利用 []int 生成 *Node
func Int2NaryNode(nodes []int) *Node {
	root := &Node{}
	if len(nodes) > 1 {
		root.Val = nodes[0]
	}
	var queue []*Node
	queue = append(queue, root)
	i := 1
	count := 0
	for i < len(nodes) {
		node := queue[0]

		var childrens []*Node
		for ; i < len(nodes) && nodes[i] != structures.NULL; i++ {
			tmp := &Node{Val: nodes[i]}
			childrens = append(childrens, tmp)
			queue = append(queue, tmp)
		}
		count++
		if count%2 == 0 {
			queue = queue[1:]
			count = 1
		}
		if node != nil {
			node.Children = childrens
		}
		i++
	}
	return root
}
