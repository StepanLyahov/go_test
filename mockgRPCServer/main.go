package main

import (
	"github.com/sirupsen/logrus"
	"go_task/mockgRPCServer/server"
)

func main() {
	_, err := server.NewServerAndStart("tcp", "localhost:8081")
	if err != nil {
		logrus.Fatalf("Error build and run GRPC-server: %v", err)
	}
}
