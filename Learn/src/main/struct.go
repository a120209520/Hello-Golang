package main

import (
	"fmt"
	"queue"
	"tree"
)

func test01() {
	fmt.Println("--------001.test01---------")
	var root tree.Node
	fmt.Println(root)
}

func test02() {
	fmt.Println("--------002.test02---------")
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{Value: 1}
	root.Right = &tree.Node{Value: 2}
	fmt.Println(root)
}

func test03() {
	fmt.Println("--------003.test03---------")
	nodes := []tree.Node{
		{Value:3},
		{},
		{1, nil, nil},
	}
	fmt.Println(nodes)
}

func test04() {
	fmt.Println("--------004.test04---------")
	node := tree.CreateNode(1)
	fmt.Println(node)
}

func test05() {
	fmt.Println("--------005.test05---------")
	node := tree.CreateNode(10)
	node.Print()
	node.SetValue(15)
	node.Print()
}

func test06() {
	fmt.Println("--------006.test06---------")
	node := tree.CreateNode(10)
	node.Left = tree.CreateNode(66)
	node.Right = tree.CreateNode(77)
	node.Print()
}

func test07() {
	fmt.Println("--------007.test07---------")
	node := tree.CreateNode(10)
	node.Left = tree.CreateNode(66)
	node.Right = tree.CreateNode(77)

	nodeEx := tree.NodeEx{node}
	nodeEx.PrintFront()
}

func test08() {
	fmt.Println("--------008.test08---------")
	q := queue.Queue{1,2,3}
	q.Push(4)
	fmt.Println(q)
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
	test08()
}