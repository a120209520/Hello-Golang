package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	fmt.Println("--------001.readFile---------")
	const filename = "src/file/abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func readFile2() {
	fmt.Println("--------002.readFile2---------")
	const filename = "src/file/abc.txt"
	//另一种写法
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func eval(a, b int, op string) int {
	fmt.Println("--------003.eval---------")
	var res int
	switch op {
	case "+": res = a + b
	case "-": res = a - b
	case "*": res = a * b
	case "/": res = a / b
	}
	return res
}

func grade(score int) string {
	fmt.Println("--------004.grade---------")
	var res string
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong input: %d", score))
	case score < 60: res = "C"
	case score < 80: res = "B"
	case score <= 100: res = "A"
	}
	return res
}

func main() {
	readFile()
	readFile2()
	fmt.Println(eval(1, 3, "+"))
	fmt.Println(eval(42, 12, "/"))
	fmt.Println(grade(78))
	fmt.Println(grade(99))
	fmt.Println(grade(-2))
}
