package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// Get 如果 Retriever 很大，则可以考虑将其设置为指针类型，防止拷贝操作过于繁重
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	resp.Body.Close()

	return string(result)
}
