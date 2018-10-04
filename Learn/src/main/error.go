package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func tryDefer() {
	//defer用于资源管理和错误处理，在任何return或panic时，会自动调用，并且是栈结构
	defer fmt.Println(1)
	fmt.Println(2)
	defer fmt.Println(3)
}

func writeFile(filename string) {
	f := fibonacci()
	file, err:= os.Create(filename)
	defer file.Close()
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i:=0; i<10; i++ {
		fmt.Fprintln(writer, f())
	}
}

func testRecover() {
	//recover仅在defer中调用，如果无法处理再重新panic
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("recover!", err)
		} else {
			panic(r)
		}
	}()
	panic(errors.New("occur error!"))
	//panic(123)
}

func main() {
	tryDefer()
	writeFile("eppl.txt")
	testRecover()
}
