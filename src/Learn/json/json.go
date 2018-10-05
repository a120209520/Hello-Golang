package json

import (
	"encoding/json"
	"fmt"
)

func basicTest() {
	fmt.Println("--------001.basicTest---------")
	//数组编码
	x := [5]int{1,2,3,4,5}
	bytes, err := json.Marshal(x)
	if err == nil {
		fmt.Println(string(bytes))
	}

	//数组解码
	var xx [5]int
	json.Unmarshal(bytes, &xx)
	fmt.Println(xx)

	//map编码
	m := make(map[string]float64)
	m["ppl"] = 1.5
	bytes, err = json.Marshal(m)
	if err == nil {
		fmt.Println(string(bytes))
	}

	//map解码
	var mm map[string]float64
	json.Unmarshal(bytes, &mm)
	fmt.Println(mm)
}

type Student struct {
	Name string  //默认以属性名作为json的key
	Age int      `json:"stuAge"`   //key重命名
	hide int     //注意，如果是小写则转换json时不可见
}

func structTest() {
	fmt.Println("--------002.structTest---------")
	//编码
	stu := Student{"ppl", 20, 123}
	bytes, err := json.Marshal(stu)
	if err == nil {
		fmt.Println(string(bytes))
	}

	//解码
	var stus Student
	json.Unmarshal(bytes, &stus)
	fmt.Println(stus)
}

func JsonTest() {
	fmt.Println("json test")
	basicTest()
	structTest()
}
