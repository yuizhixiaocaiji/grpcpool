package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcpool/greeter_grpc/proto"
	"grpcpool/grpc_client_pool"
	"log"
)

var (
	addr = flag.String("addr", "localhost:50051", "")
)

func main() {
	flag.Parse()
	/*
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		sayHello(conn)
	*/
	pool, err := grpc_client_pool.GetPool(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	conn := pool.Get()
	sayHello(conn)
	defer conn.Close()
	pool.Put(conn)
}

func sayHello(conn *grpc.ClientConn) {
	c := proto.NewGreeterClient(conn)
	ctx := context.Background()
	in := &proto.HelloRequest{
		Msg: "Hello Server",
	}
	r, err := c.SayHello(ctx, in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client Recv: ", r.Msg)
}
