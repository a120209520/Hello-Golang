package gochat

import "time"

func Main() {
	go ServerInit()
	go ClientInit()

	time.Sleep(time.Second*120)
}
