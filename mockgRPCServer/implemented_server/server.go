package implemented_server

import (
	"context"
	"go_task/internal/adapter/delivery/rpc"
)

type GRPCServer struct {
}

func (G GRPCServer) Calculate(ctx context.Context, request *rpc.ParamsRequest) (*rpc.CalculateResponse, error) {
	return &rpc.CalculateResponse{Uuid: "test", Err: ""}, nil
}
