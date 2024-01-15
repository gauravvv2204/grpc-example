package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "main/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultKey = "world"
	defaultVal = "is round"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	key  = flag.String("key", defaultKey, "it is key")
	val  = flag.String("val", defaultVal, "value")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreetClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Key: *key, Val: *val})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", r.GetVal())
}

