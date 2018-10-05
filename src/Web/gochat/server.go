package gochat

import (
	"fmt"
	"net"
	"strings"
)

var userConnMap map[string]net.Conn
var msgQueue chan string
var quitChan chan bool

func MsgProcesser(msg string) {
	strs := strings.Split(msg, "#")
	if len(strs) > 1 {
		addr := strs[0]
		content := strs[1]
		if connect, ok := userConnMap[addr]; ok {
			_, err := connect.Write([]byte(content))
			CheckError(err)
		}
	} else {

	}
}

func MsgDispatcher() {
	for  {
		select {
		case msg := <- msgQueue:
			go MsgProcesser(msg)
		case <- quitChan:
			break
		}
	}
}

func ServerTask(connect net.Conn) {
	buf := make([]byte, 1024)
	defer connect.Close()
	for {
		n, err := connect.Read(buf)
		if err == nil {
			msg := string(buf[:n])
			msgQueue <- msg
		}
	}
}

func PrintUserMap() {
	fmt.Println("----------------current userMap-----------------")
	for v := range userConnMap {
		fmt.Println("[userMap]:", v)
	}
}

func ServerInit() {
	listen, err := net.Listen("tcp", SERVER_HOST)
	defer listen.Close()
	CheckError(err)

	userConnMap = make(map[string]net.Conn)
	msgQueue = make(chan string, 1000)
	quitChan = make(chan bool)


	go MsgDispatcher()
	for {
		connect, err := listen.Accept()
		CheckError(err)
		userConnMap[connect.RemoteAddr().String()] = connect
		PrintUserMap()
		go ServerTask(connect)
	}
}
