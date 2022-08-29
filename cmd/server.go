package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/beanyoung/dddgo/app"
	"github.com/beanyoung/dddgo/rpc/order"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	order.RegisterOrderServer(grpcServer, app.NewOrderServer())
	log.Println("start to serve")
	grpcServer.Serve(lis)
}
