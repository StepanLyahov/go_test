package calculator

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"go_task/internal/adapter/delivery/rpc"
)

type Calculator struct {
}

func (G Calculator) Calculate(ctx context.Context, request *rpc.ParamsRequest) (*rpc.CalculateResponse, error) {
	if request == nil {
		return &rpc.CalculateResponse{Uuid: "", Err: ""}, errors.New("request is nil")
	}

	s, err := expect(request.Param1, request.Param2)
	if err != nil {
		logrus.Infof("Calculate err: %v", err)
		return &rpc.CalculateResponse{Uuid: "", Err: err.Error()}, nil
	}

	return &rpc.CalculateResponse{
			Uuid: s,
			Err:  ""},
		nil
}

func expect(param1, param2 string) (string, error) {

	if len(param1) == 0 || len(param2) == 0 {
		return "", errors.New("len params must be not empty")
	}

	if len(param1) != len(param2) {
		return "", errors.New("len params must be equal")
	}

	dataParam1 := []byte(param1)
	dataParam2 := []byte(param2)

	res := make([]byte, 0)
	for i, _ := range dataParam1 {
		res = append(res, dataParam1[i]&dataParam2[i])
	}

	return string(res), nil
}
