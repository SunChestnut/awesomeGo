package main

import (
	pb "awesomeGo/rpcdemo/grpcdemo/helloworld"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:9000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

// callSayHello ==> 调用 Server 端的 SayHello 函数
func callSayHello(client pb.GreeterClient, parent context.Context) {
	ctx, cancelFunc := context.WithTimeout(parent, 10*time.Second)
	defer cancelFunc()

	reply, err := client.SayHello(ctx, &pb.HelloRequest{
		Id:   1,
		Name: *name,
	})
	if err != nil {
		log.Fatalf("💔could not greet: %v", err)
	}
	log.Printf("❤️Greeting %s on %v\n", reply.GetMessage(), reply.GetReplyTime())
}

// callHealthCheck ==> 调用 Server 端的 HealthCheck 函数
func callHealthCheck(client pb.GreeterClient, parent context.Context) {
	ctx, cancelFunc := context.WithTimeout(parent, 10*time.Second)
	defer cancelFunc()

	checkReply, err := client.HealthCheck(ctx, &pb.Check{Id: 2})
	if err != nil {
		log.Fatalf("💔could not check health: %v", err)
	}
	log.Printf("❤️Success to check health: %s\n", checkReply.Message)
}

func main() {
	flag.Parse()

	// Set up a connection to the greeter_server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	parent := context.Background()

	for {
		//// Contact the greeter_server and print out its response.
		//ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)

		callSayHello(client, parent)
		time.Sleep(2 * time.Second)

		callHealthCheck(client, parent)
		time.Sleep(2 * time.Second)
	}

}
