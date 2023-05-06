package main

import (
	pb "awesomeGo/rpcdemo/grpcdemo/helloworld"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var (
	port = flag.Int("port", 9000, "The server listens on the port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("🍏Received: %v", in.GetName())
	return &pb.HelloReply{
		Message:   "Hello" + in.GetName(),
		ReplyTime: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *server) HealthCheck(ctx context.Context, check *pb.Check) (*pb.CheckReply, error) {
	log.Printf("🍏Received: %v", check.GetId())
	return &pb.CheckReply{Message: "fit as a fiddle"}, nil
}

func main() {
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatalf("fail to lisent: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("fail to server: %v", err)
	}
}
