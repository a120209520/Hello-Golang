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

type Point struct {
	x, y int
}

func (p *Point) set(x, y int) {
	p.x = x
	p.y = y
}

func (p *Point) get() (int, int) {
	return p.x, p.y
}

type Point3D struct {
	//继承
	Point
	z int
}

func (p *Point3D) get3D() (int, int, int) {
	return p.x, p.y, p.z //子类是可以直接使用父类的函数和域的
}

func (p *Point3D) get2D() (int, int) {
	return p.get()
}

func test09() {
	fmt.Println("--------009.test09---------")
	var p Point
	fmt.Printf("%T\n", p)  //值类型
	//var p *Point = &Point{}
	//fmt.Printf("%T\n", p)  //指针类型
	p.set(1, 2)         //调用函数时，会自动判断值类型还是指针类型
	x, y := p.get()
	fmt.Println(x, y)
}

func test10() {
	fmt.Println("--------010.test10---------")
	var p Point3D
	x, y, z := 0, 0, 0
	p.set(1, 2)
	x, y = p.get2D()
	x, y, z = p.get3D()
	fmt.Println(p)
	fmt.Println(x, y, z)
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
	test09()
	test10()
}