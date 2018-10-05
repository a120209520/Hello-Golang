package main

import (
	"fmt"
)

func array() {
	fmt.Println("--------001.array---------")
	//数组的定义
	var arr1 [5]int         //若不指定，则初始化为0
	arr2 := [3]int{2, 4, 5} //:=必须初始化
	arr3 := [...]int{3, 43, 5, 6, 6, 7}
	var arr4 [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(arr4)

	//数组的遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("%d-",arr3[i])
	}
	fmt.Println()
	for i := range arr3 {
		fmt.Printf("%d-",arr3[i])
	}
	fmt.Println()
	for _,v := range arr3 {
		fmt.Printf("%d-",v)
	}
	fmt.Println()
}

func printArray(arr [5]int) {
	for _,v := range arr {
		fmt.Printf("%d-",v)
	}
	fmt.Println()
}

func editArray(arr [5]int) {
	//数组是值类型，修改无效
	arr[0] = 100
}

func editArray2(arr *[5]int) {
	//数组就是一个特殊类型，和int一样取指针才能修改
	arr[0] = 100      //注意，对于数组指针或者数组来说，都用arr[i]来访问，有点怪异
}

func main() {
	array()
	arr := [5]int{34,4,5,5,6}
	printArray(arr)
	editArray(arr)
	printArray(arr)
	editArray2(&arr)
	printArray(arr)
}