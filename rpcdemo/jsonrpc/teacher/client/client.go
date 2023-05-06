package main

import (
	rpcdemo "awesomeGo/rpcdemo/jsonrpc/teacher"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	if err := client.Call("DemoService.Div", rpcdemo.Args{A: 99, B: 8}, &result); err != nil {
		log.Fatalf("error call method: %v", err)
	} else {
		log.Println(result)
	}

	if err := client.Call("DemoService.Div", rpcdemo.Args{A: 6666, B: 0}, &result); err != nil {
		log.Fatalf("error call method: %v", err)
	} else {
		log.Println(result)
	}
}
