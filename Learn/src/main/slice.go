package main

import "fmt"

var arr = [...]int{0,1,2,3,4,5,6,7,8,9}

func slice() {
	fmt.Println("--------001.slice---------")
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])
}

func editSlice(arr []int) {
	fmt.Println("--------002.editSlice---------")
	//slice是对arr的一个视图，会实际修改slice
	arr[0] = 100
}

func reSlice() {
	fmt.Println("--------003.reSlice---------")
	s := arr[:5]
	s2 := s[2:]
	fmt.Println(s2)
}

func extendSlice() {
	fmt.Println("--------004.extendSlice---------")
	s := arr[2:6]
	es := s[3:5]
	//es2 := s[-1:2]    //error, 不允许向左扩展取值
	es3 := s[3:]        //默认不扩展
	fmt.Println(s)
	fmt.Println(es)     //可以向右扩展取值
	fmt.Println(es3)
	//fmt.Println(s[4]) //但是不可以直接取slice的扩展值，out of range
	fmt.Printf("s=%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
	fmt.Printf("es=%v, len(es)=%d, cap(es)=%d\n", es, len(es), cap(es))
	fmt.Printf("es3=%v, len(es3)=%d, cap(es3)=%d\n", es3, len(es3), cap(es3))
	fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))
}

func appendSlice() {
	fmt.Println("--------005.appendSlice---------")
	s := arr[2:]
	s2 := append(s, 66) //此时s2和s3不再是arr的切片，而是重新分配了新内存
	s3 := append(s2, 77)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Printf("s=%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("arr=%v, len(arr)=%d, cap(arr)=%d\n", arr, len(arr), cap(arr))
}

func createSlice() {
	fmt.Println("--------006.createSlice---------")
	var s []int
	fmt.Println(s)
	for i:=0; i<10; i++ {
		s = append(s, i)
		fmt.Printf("s=%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
	}

	s1 := make([]int, 10)
	s2 := make([]int, 10, 32)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}

func copySlice() {
	fmt.Println("--------007.copySlice---------")
	s := make([]int, 10)
	s1 := arr[:3]
	copy(s, s1)
	fmt.Println(s)
}

func deleteSlice() {
	fmt.Println("--------008.deleteSlice---------")
	s := arr[2:6]
	s1 := append(s[:1], s[2:]...)
	fmt.Println(s1)
}

func main() {
	slice()
	arr := [...]int{10,11,12,13,14,15,16,17,18,19}
	s1 := arr[2:]
	s2 := arr[:]
	editSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	editSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
	reSlice()
	extendSlice()
	appendSlice()
	createSlice()
	copySlice()
	deleteSlice()
}

