package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func str1() {
	s := "你是皮皮落"
	fmt.Println(len(s))  //字节数
	fmt.Println(utf8.RuneCountInString(s))  //字符数
	for _, ch := range s {
		fmt.Printf("(%c)", ch)
	}
	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
}

func main() {
	str1()
}
