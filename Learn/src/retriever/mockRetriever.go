package retriever

import "fmt"

type MockRetriever struct {
	Contents string
}

func (r *MockRetriever) Get(url string) string {
	return r.Contents
}

func (r *MockRetriever) Post(url string, form map[string]string) string {
	r.Contents = form["course"]
	return r.Contents
}

func (r *MockRetriever) String() string {
	return fmt.Sprintf("{content}=%s", r.Contents)
}








