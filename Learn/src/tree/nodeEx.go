package tree

import "fmt"

type NodeEx struct {
	Node *Node
}

func (node *NodeEx) PrintFront() {
	if node == nil || node.Node == nil {
		return
	}
	left := NodeEx{node.Node.Left}
	left.PrintFront()
	right := NodeEx{node.Node.Right}
	right.PrintFront()
	fmt.Println(node.Node.Value)
}