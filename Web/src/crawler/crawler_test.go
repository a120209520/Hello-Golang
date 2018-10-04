package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCrawler(t *testing.T) {
	url := "https://book.douban.com/subject/30193594/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(html))
}
