package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		c1 <- 1
	}()
	go func() {
		c2 <- 2
	}()
	time.Sleep(time.Second)
	select {
	case n := <- c1:
		fmt.Println("c1=", n)
	case n := <- c2:
		fmt.Println("c2=", n)
	default:
		fmt.Println("no value")
	}
}
