package main

import (
	"fmt"
	"sync"
	"time"
)

/*func recver(c chan int) {
	for {
		data, ok := <- c
		if !ok {
			break
		}
		fmt.Println(data)
	}
	fmt.Println("finish!")
}*/
func recver(c chan int) {
	for n := range c {
		fmt.Printf("data=%d\n", n)
	}
	fmt.Println("finish!")
}

func recver2(c chan int, d chan bool) {
	for n := range c {
		fmt.Printf("data=%d\n", n)
	}
	fmt.Println("finish!")
	time.Sleep(5*time.Second)
	d <- true
}

func recver3(c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("data=%d\n", n)
	}
	fmt.Println("finish!")
	time.Sleep(5*time.Second)
	wg.Done()
}

func chanTest() {
	var c chan int
	c = make(chan int)  //发送后必须有其他协程收，否则会一直阻塞
	go recver(c)
	c <- 1
	c <- 2
	close(c)
}

func bufChanTest() {
	var c chan int
	c = make(chan int, 3)  //缓冲区为3，未满时不会阻塞
	go recver(c)
	c <- 1
	time.Sleep(time.Second)
	c <- 2
	time.Sleep(time.Second)
	c <- 3
	time.Sleep(time.Second)
	c <- 4
	time.Sleep(time.Second)
	c <- 0
	close(c)  //close后，接收方会收到0
}

func doneTest() {
	c := make(chan int, 3)
	d := make(chan bool)
	go recver2(c, d)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 0
	close(c)

	<- d  //再次阻塞等待完成消息
	fmt.Println("is done!")
}

func waitGroupTest() {
	c := make(chan int)
	var wg sync.WaitGroup
	go recver3(c, &wg)
	wg.Add(1)  //表示有多少个channel
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 0
	close(c)

	wg.Wait()  //再次阻塞等待完成消息
	fmt.Println("is done!")
}


func main() {
	//chanTest()
	//bufChanTest()
	//doneTest()
	waitGroupTest()
	time.Sleep(time.Second)
}