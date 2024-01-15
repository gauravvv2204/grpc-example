package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "main/proto"

	redis "github.com/redis/go-redis/v9"
	grpc "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server Port")

	client = redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

type server struct {
	pb.UnimplementedGreetServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: key: %v val: %v", in.GetKey(), in.GetVal())

	err := client.Set(ctx, in.GetKey(), in.GetVal(), 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, in.GetKey()).Result()
	if err != nil {
		panic(err)
	}

	return &pb.HelloReply{Val: "Val: " + val}, nil
}

func main() {
	flag.Parse()
	//ctx := context.Background()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	s.Serve(lis)
}
