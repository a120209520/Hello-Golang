package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) incr() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	return a.value
}

func main() {
	var a atomicInt
	for i:=0; i<100; i++ {
		go func() {a.incr()}()
	}
	time.Sleep(time.Second)
	fmt.Println(a.value)
}