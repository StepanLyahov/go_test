package main

import (
	"github.com/sirupsen/logrus"
	"go_task/internal/adapter/delivery/rpc"
	"go_task/mockgRPCServer/implemented_server"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	impl := &implemented_server.GRPCServer{}

	rpc.RegisterCalculatorUUIDServiceServer(server, impl)

	l, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		logrus.Fatal(err)
	}

	if err := server.Serve(l); err != nil {
		logrus.Fatal(err)
	}
}
