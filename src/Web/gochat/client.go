package gochat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func MessageSend(connect net.Conn, id int) {
	var input string
	defer connect.Close()
	for  {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)
		if strings.ToUpper(input) == "EXIT" {
			break
		}
		_, err := connect.Write(data)
		CheckError(err)
		fmt.Printf("[client%d][send %d]=%s\n", id, len(data), string(data))
	}
}

func MessageRecv(connect net.Conn, id int) {
	buf := make([]byte, 1024)
	defer connect.Close()
	for {
		n, err := connect.Read(buf)
		if err == nil {
			fmt.Printf("[client%d][recv %d]=%s\n", id, n, string(buf[:n]))
		}
	}
}

func ClientTask(id int) {
	connect, err := net.Dial("tcp", SERVER_HOST)
	CheckError(err)
	go MessageSend(connect, id)
	go MessageRecv(connect, id)
}

func ClientInit() {
	go ClientTask(1)
	go ClientTask(2)
}
