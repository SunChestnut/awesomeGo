package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// 访问手机版的 IMOOC
func getWebSite() {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com/", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36")

	// 重新实现 http.Client，以便查看网站是否发生了重定向。
	// 如果发生了重定向，其地址存储在 via 中；其结果中，如果 error 值为 nil，则表示让 greeter_client 正常去重定向，如果返回错误则会终止重定向
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			//fmt.Printf("via : %v\n", via)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}

func main() {
	//getWebSite()
}
