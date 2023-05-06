package main

import (
	"awesomeGo/rpc/jsonrpc/syl/server/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
*
* JSON-RPC Server 端
🔗https://www.jsonrpc.org/specification
*
* 单独启动服务端，可在控制台访问：
*
* @author  sun
* @date 2022/11/8 16:08
*/

func main() {
	if err := rpc.Register(service.WorkerService{}); err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("🍑Start RPC Server on port %v\n", ":1234")

	for {
		if conn, err := listen.Accept(); err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		} else {
			// 如果不使用协程单独接收 RPC Client 端的请求，下述代码每次接收并处理一个请求
			go jsonrpc.ServeConn(conn)
		}

	}

}
