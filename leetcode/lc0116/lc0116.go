package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	helpConnect(root.Left, root.Right)
	return root
}

func helpConnect(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	helpConnect(node1.Left, node1.Right)
	helpConnect(node2.Left, node2.Right)
	helpConnect(node1.Right, node2.Left)
}
