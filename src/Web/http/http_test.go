package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDemo(t *testing.T) {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello golang!!"))
	})

	http.ListenAndServe("127.0.0.1:8080", nil)   //阻塞
	fmt.Println("server shutdown")
}
