package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	client := &http.Client{}

	// https://www.zhihu.com/hot
	req, err := http.NewRequest("GET", "https://movie.douban.com/", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "ABC")
	fmt.Printf("User-Agent: %v\n", req.Header.Get("User-Agent"))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Printf("resp-status: %v\n", resp.Status)

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.OpenFile("./request/index.html", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err = fmt.Fprintln(writer, string(bytes)); err != nil {
		log.Fatalln(err)
	}
}
