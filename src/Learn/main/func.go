package main

import "fmt"

func div(a, b int) (q, r int) {
	return a/b, a%b
}

func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}

func main() {
	q, r := div(5, 2)
	fmt.Println(q, r)

	q1, _ := div(5, 2)  //不想使用的参数用_代替
	fmt.Println(q1)

	q2, r2 := div2(9, 2)
	fmt.Println(q2, r2)

	fmt.Println(sum(2,4,5))
}