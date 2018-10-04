package main

import "fmt"

//函数式编程
//函数的参数、返回值都可以是函数

//正统函数式编程，不能有变量，只有常量和函数————一般不会严格遵守这样的规定

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func adderTest() {
	a := adder()  //a可以当作是函数指针
	fmt.Println(a(10))
	fmt.Println(a(11))
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fibonacciTest() {
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func main() {
	adderTest()
	fibonacciTest()
}