package main

import (
	"fmt"
)

type Animal interface {
	Fly()
	Run()
}

type FlyAnimal interface {
	Fly()
}

type Bird struct {

}

func (b Bird) Fly() {
	fmt.Println("bird is flying")
}

func (b Bird) Run() {
	fmt.Println("bird is runing")
}

func interface2Test() {
	fmt.Println("--------001.interface2Test---------")
	var a Animal
	var fa FlyAnimal
	b := new(Bird)
	a = b
	a.Fly()
	a.Run()

	fa = b     //类可以赋值给接口
	fa.Fly()

	fa = a     //大接口也可以赋值给小接口
	fa.Fly()
}

func interface2Test2() {
	fmt.Println("--------002.interface2Test2---------")
	var allType interface{}
	allType = "string"
	fmt.Println(allType)
	fmt.Printf("%T\n", allType)
	_, ok := allType.(string)
	fmt.Printf("allType is string ? %v\n", ok)
	allType = 213
	fmt.Println(allType)
	fmt.Printf("%T\n", allType)
	allType = new(Bird)
	fmt.Println(allType)
	fmt.Printf("%T\n", allType)
}

func interface2Test3() {
	fmt.Println("--------003.interface2Test3---------")
	var allType interface{}
	//allType = "string"
	//allType = 1.4523
	allType = 1542543254          //默认类型就是int, 不是任何int8, int16...等等
	switch allType.(type) {
	/*case int:
		fmt.Println("this is a int")*/
	case int8:
		fmt.Println("this is a int8")
	case int16:
		fmt.Println("this is a int16")
	case int32:
		fmt.Println("this is a int32")
	case int64:
		fmt.Println("this is a int64")
	case string:
		fmt.Println("this is a string")
	case float64:
		fmt.Println("this is a float64")
	case float32:
		fmt.Println("this is a float32")
	}
}

func main() {
	interface2Test()
	interface2Test2()
	interface2Test3()
}