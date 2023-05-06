package main

import (
	rpcdemo "awesomeGo/rpcdemo/jsonrpc/teacher"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpcdemo.DemoService{})
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error : %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
