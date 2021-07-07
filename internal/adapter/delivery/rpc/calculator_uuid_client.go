package rpc

import (
	"context"
	"errors"
	"fmt"
	"go_task/internal/config"
	"google.golang.org/grpc"
)

type CalculatorUuidClient struct {
	client CalculatorUUIDServiceClient
}

func NewCalculatorUuidClient(cnf *config.Config) (*CalculatorUuidClient, error) {
	target := fmt.Sprintf("%v:%v", cnf.GRPC.Host, cnf.GRPC.Port)

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := NewCalculatorUUIDServiceClient(conn)

	return &CalculatorUuidClient{client}, nil
}

func (c *CalculatorUuidClient) Execute(param1, param2 string) (string, error) {
	in := ParamsRequest{Param1: param1, Param2: param2}

	calculate, err := c.client.Calculate(context.Background(), &in)
	if err != nil {
		return "", err
	}

	if len(calculate.Err) != 0 {
		return "", errors.New(calculate.Err)
	}

	return calculate.Uuid, nil
}
