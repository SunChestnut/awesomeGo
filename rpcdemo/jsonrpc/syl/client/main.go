package main

import (
	"awesomeGo/rpcdemo/jsonrpc/syl/server/service"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

/**
*
* TODO
*
* @author  sun
* @date 2022/11/8 16:08
 */

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	client := jsonrpc.NewClient(conn)

	result := service.Result{}
	err = client.Call("WorkerService.HealthCheck", "Hello~", &result)
	if err != nil {
		log.Fatal(err)
	}
	marshal, err := json.Marshal(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(marshal))

	var ans service.Result
	err = client.Call("WorkerService.ClimbStairs",
		service.Request{
			Msg: "req from syl",
			Num: 10,
		},
		&ans)
	if err != nil {
		log.Fatal(err)
	}

}
