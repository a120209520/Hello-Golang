package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for ; n > 0; n/=2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

//一运行就死机。。Goland的处理机制有点问题啊
func forever() {
	for {
		fmt.Printf("abab")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(0))
	forever()
}
