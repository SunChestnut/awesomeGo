package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s ❤️", r.URL.Path[1:])
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "read the data: %s\n", string(bytes))

	bytes, err = io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "read the data: %s\n", string(bytes))
}

func readMulti(w http.ResponseWriter, r *http.Request) {
	readCloser, err := r.GetBody()
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(readCloser)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "read the data: %s\n", string(bytes))
	//
	//body = r.GetBody
	//readCloser, err = body()
	//if err != nil {
	//	panic(err)
	//}
	//bytes, err = io.ReadAll(readCloser)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Fprintf(w, "read the data: %s\n", string(bytes))

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/sayHello", sayHello)
	http.HandleFunc("/readBodyOnce", readBodyOnce)
	http.HandleFunc("/readMulti", readMulti)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
