package server

import (
	"go_task/internal/adapter/delivery/rpc"
	"go_task/mockgRPCServer/calculator"
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	server   *grpc.Server
	listener *net.Listener
}

func NewServerAndStart(network, address string) (GRPCServer, error) {
	server := grpc.NewServer()
	impl := &calculator.Calculator{}

	rpc.RegisterCalculatorUUIDServiceServer(server, impl)

	l, err := net.Listen(network, address)
	if err != nil {
		return GRPCServer{}, err
	}

	if err := server.Serve(l); err != nil {
		return GRPCServer{}, err
	}

	return GRPCServer{
		server:   server,
		listener: &l,
	}, nil
}
