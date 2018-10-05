package main

import "fmt"
//go只有值传递，无效
func swap(a, b int) {
	fmt.Println("--------001.swap---------")
	b, a = a, b
}

func swap2(a, b *int) {
	fmt.Println("--------002.swap2---------")
	*b, *a = *a, *b
}

func swap3(a, b int) (int, int){
	fmt.Println("--------003.swap3---------")
	return b, a
}

func main() {
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)
	swap2(&a, &b)
	fmt.Println(a, b)
	a, b = swap3(a, b)
	fmt.Println(a, b)
}