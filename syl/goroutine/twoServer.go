package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello you~")
	})

	err := http.ListenAndServe("", mux)
	if err != nil {
		log.Fatal(err)
	}

	go http.ListenAndServe("", http.DefaultServeMux)
}
