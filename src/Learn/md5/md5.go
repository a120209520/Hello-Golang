package md5

import (
	"crypto/md5"
	"fmt"
)

func Md5Test() {
	md5i := md5.New()
	md5i.Write([]byte("123321"))
	res := md5i.Sum([]byte(""))
	fmt.Printf("%x\n", res)
}