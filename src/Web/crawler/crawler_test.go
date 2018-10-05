package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

func TestCrawler(t *testing.T) {
	//获取html内容
	url := "https://book.douban.com/subject/30193594/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, _ := ioutil.ReadAll(resp.Body)
	htmlString := string(html)
	//fmt.Println(string(html))

	reg := regexp.MustCompile(`<strong\s*class="ll rating_num\s*"\s*property="v:average">(.*)</strong>`)
	res := reg.FindAllString(htmlString, -1)
	res2 := reg.FindAllStringSubmatch(htmlString, -1)
	fmt.Println(res)
	fmt.Println(res2[0][1])
}
