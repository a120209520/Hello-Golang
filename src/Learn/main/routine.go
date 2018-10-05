package main

import (
	"fmt"
	"time"
)

var a int = 0

func task(i int) {
	fmt.Printf("id=%d\n", i)
}

func add() {
	if(a < 100) {
		time.Sleep(10)
		a++
	}
}

func routineTest() {
	for i:=0; i<1000; i++ {
		go task(i)
	}
}

func routineTest2() {
	for i:=0; i<1000; i++ {
		go add()
	}
	fmt.Println(a)  //如何解决并发？
}

func main() {
	//routineTest()
	routineTest2()
	time.Sleep(100)
}
