package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//自定义构造函数
func CreateNode(v int) *Node {
	return &Node{Value: v}
}

//结构体方法(实际上只是定义了接收者的普通方法)
//调用方法: Node.print()
func (node *Node) Print() {
	if node == nil {
		return
	}
	node.Left.Print()
	fmt.Println(node.Value)
	node.Right.Print()
}
func (node *Node) SetValue(v int) {
	node.Value = v  //和数组一样，这里node也不用写成*Node.Value，不统一，觉得很怪异。。
}

