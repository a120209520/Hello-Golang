package main

import (
	"fmt"
	"retriever"
	"time"
)
//接口可以表示一个值或者指针
type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}
type RetrieverPoster interface {
	Retriever
	Poster
}


func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(p Poster) {
	p.Post("http://www.imooc.com",
		map[string]string{
			"name":"aaa",
			"course":"golang",
		})
}

func session(s RetrieverPoster) {
	s.Post("http://www.imooc.com",
		map[string]string{
			"name":"aaa",
			"course":"golang",
		})
	fmt.Println(s.Get("http://www.imooc.com"))
}

func interfaceTest() {
	fmt.Println("--------001.interfaceTest---------")
	r := retriever.MockRetriever{"hello ppl"}
	fmt.Println(download(&r))
}

func interfaceTest2() {
	fmt.Println("--------002.interfaceTest2---------")
	r := retriever.RealRetriever{}
	fmt.Println(download(&r))
}

func interfaceTest3() {
	fmt.Println("--------003.interfaceTest3---------")
	var rp Retriever
	rp = &retriever.MockRetriever{"hello"}
	inspect(rp)
	rp = &retriever.RealRetriever {
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	inspect(rp)

	//Type Assertion
	rr := rp.(*retriever.RealRetriever)
	fmt.Println(rr.UserAgent)
}

func interfaceTest4() {
	fmt.Println("--------004.interfaceTest4---------")
	var rp RetrieverPoster
	rp = &retriever.MockRetriever{}
	session(rp)
}




func interfaceStandand() {
	fmt.Println("--------005.interfaceStandand---------")

}

func inspect(r Retriever) {
	//Type Switch
	switch v := r.(type) {
	case *retriever.MockRetriever:
		fmt.Printf("mock: %T %v\n", v, v)  //如果实现了stringer接口, %v会自动调用String()函数
	case *retriever.RealRetriever:
		fmt.Printf("real: %T %v\n", v, v)
	}
}

func main() {
	interfaceTest()
	interfaceTest2()
	interfaceTest3()
	interfaceTest4()
}