package retriever

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type RealRetriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r *RealRetriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	res, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}
	return string(res)
}


