package main

import (
	"fmt"
)

func hello()  {
	fmt.Println("--------001.hello---------")
	fmt.Println("hello go!")
}

func variable() {
	fmt.Println("--------002.variable---------")
	var a int
	var s string
	fmt.Printf("%d %s\n", a, s)
	fmt.Printf("%d %q\n", a, s)   //%q是带引号的字符串
}

func variableInit() {
	fmt.Println("--------003.variableInit---------")
	var a int = 3
	var b = 4  //可以省略类型
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableAuto()  {
	fmt.Println("--------004.variableAuto---------")
	var a, b, c  = 1, 2, "xyz"
	fmt.Println(a, b, c)
}

func variableShorter()  {
	fmt.Println("--------005.variableShorter---------")
	a, b, c  := 1, 2, "xyz"
	fmt.Println(a, b, c)
}

var ga = 3
var gb = true
//gb := "gb" //包变量不能用:=
var (
	gc = 123
	gd = "hello"
)

func global()  {
	fmt.Println("--------005.global---------")
	fmt.Println(ga, gb, gc, gd)
}

func main() {
	hello()
	variable()
	variableInit()
	variableAuto()
	variableShorter()
	global()
}