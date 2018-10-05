package gochat

import (
	"fmt"
)

const SERVER_HOST = "127.0.0.1:8088"

func CheckError(err error) {
	if err != nil {
		fmt.Println("[ERROR]:", err.Error())
	}
}