package main

import (
	"fmt"
	"github.com/astaxie/goredis"
	"testing"
)

func set(cli *goredis.Client) {
	err := cli.Set("go", []byte("hello, golang"))
	if err != nil {
		fmt.Println("set redis error!")
	} else {
		fmt.Println("set redis success!")
	}
}

func get(cli *goredis.Client) {
	res, err := cli.Get("go")
	if err != nil {
		fmt.Println("get redis error!")
	} else {
		fmt.Println("get redis success!")
		fmt.Println(string(res))
	}
}

func hmset(cli *goredis.Client) {
	m := make(map[string]interface{})
	m["A"] = "bbccaa"
	m["B"] = 123
	err := cli.Hmset("hash", m)
	if err != nil {
		fmt.Println("hmset redis error!")
	} else {
		fmt.Println("hmset redis success!")
	}
}

func TestRedis(t *testing.T) {
	var cli goredis.Client
	cli.Addr = "192.168.25.133:6379"
	//set(&cli)
	//get(&cli)
	hmset(&cli)
}

