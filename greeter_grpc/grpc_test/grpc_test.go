package grpc_test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcpool/greeter_grpc/proto"
	"grpcpool/grpc_client_pool"
	"log"
	"testing"
)

func BenchmarkGrpcWithOutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			b.Error(err)
		}
		in := &proto.HelloRequest{
			Msg: "Hello server",
		}
		c := proto.NewGreeterClient(conn)
		c.SayHello(context.Background(), in)
	}
}

func BenchmarkGrpcWithPool(b *testing.B) {
	pool, err := grpc_client_pool.GetPool("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		conn := pool.Get()
		in := &proto.HelloRequest{
			Msg: "Hello server",
		}
		c := proto.NewGreeterClient(conn)
		c.SayHello(context.Background(), in)
		pool.Put(conn)
	}
}
