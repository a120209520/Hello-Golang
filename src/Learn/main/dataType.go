package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

func intType() {
	fmt.Println("--------001.intType---------")
	var a int = 3  //根据系统或编译器，默认int32或者int64
	var b int8 = -128  //[-128, 127]
	c := 666       //根据系统或编译器，默认int32或者int64
	var d uint = 999
	fmt.Println(a, b, c, d)
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c), unsafe.Sizeof(d))
}

func runeType() {
	fmt.Println("--------002.runeType---------")
	var a rune = 'a'    //长度32
	var b rune = '啊'
	fmt.Printf("%c %c\n", a, b)
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}

func floatType() {
	fmt.Println("--------003.floatType---------")
	var a float32 = 1.1
	var b float64 = 2.2
	c := a     //这种情况赋值，c和a保持同一个类型
	d := 3.3   //默认是float64
	fmt.Println(a, b, c, d)
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c), unsafe.Sizeof(d))
}

func complexType() {
	fmt.Println("--------004.complexType---------")
	var c complex128 = 3 + 4i
	d := 3 + 4i  //默认complex128
	res := cmplx.Abs(c)
	res2 := cmplx.Abs(d)
	res3 := cmplx.Pow(math.E, 1i*math.Pi)
	fmt.Println(res)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(unsafe.Sizeof(c), unsafe.Sizeof(d))
}

func typeTrans() {
	fmt.Println("--------005.typeTrans---------")
	a := 3
	b := 4
	var c int = int(math.Sqrt(float64(a*a+b*b)))  //必须显式类型转换
	fmt.Println(c)
}

func constType() {
	fmt.Println("--------006.constType---------")
	const size int = 100
	//不指定类型的常量类似于宏，类型是动态的
	const filename = "aaaaaaa"
	const a, b = 3, 4
	c := math.Sqrt(a*a + b*b) //这里是float
	var d int = a                //这里变成int
	fmt.Println(c, d)
	fmt.Println(filename)
	fmt.Println(size)
}

func enumType() {
	fmt.Println("--------007.enumType---------")
	//模拟枚举类型
	const (
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)

	const (
		aa = 3
		bb = iota  //表示自增，从0开始
		_         //表示跳过
		cc
		dd
	)

	const (
		b = 1<<(10*iota)
		kb
		mb
		gb
	)
	fmt.Println(python)
	fmt.Println(aa, bb, cc, dd)
	fmt.Println(b, kb, mb, gb)
}

func main() {
	intType()
	runeType()
	floatType()
	complexType()
	typeTrans()
	constType()
	enumType()
}
