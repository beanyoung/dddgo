package main

import (
	"context"
	"flag"
	"log"

	"github.com/beanyoung/dddgo/rpc/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := order.NewOrderClient(conn)
	resp, err := client.Consume(context.Background(), &order.ConsumeRequest{})
	log.Println(resp, err)
}
