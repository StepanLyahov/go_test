package calculator

import (
	"context"
	"github.com/sirupsen/logrus"
	"go_task/internal/adapter/delivery/rpc"
	"testing"
)

func TestExpect_ValidParams(t *testing.T) {
	str1 := "abc"
	str2 := "dej"

	s, err := expect(str1, str2)
	if err != nil {
		t.Fatalf("Err must be null, but err: %v", err)
	}

	logrus.Info(s)
}

func TestExpect_InvalidParams(t *testing.T) {
	str1 := "abc"
	str2 := "de"

	_, err := expect(str1, str2)
	if err == nil {
		t.Fatalf("Err must be 'len params must be equal' ")
	}

	str1 = ""
	str2 = ""

	_, err = expect(str1, str2)
	if err == nil {
		t.Fatalf("Err must be 'len params must be equal' ")
	}
}

func TestExpect_Symmetry(t *testing.T) {
	str1 := "abc"
	str2 := "dej"

	res1, _ := expect(str1, str2)
	res2, _ := expect(str2, str1)
	if res1 != res2 {
		t.Fatalf("Alg must be symmetry")
	}

}

func TestCalculator_Calculate_Valid(t *testing.T) {
	calculator := Calculator{}

	request := rpc.ParamsRequest{
		Param1: "test1",
		Param2: "test2",
	}

	calculate, err := calculator.Calculate(context.Background(), &request)
	if err != nil {
		t.Fatalf("Err must be nil, but err: %v", err)
	}

	if len(calculate.Err) != 0 {
		t.Fatalf("Err in responce must be nil, but err: %v", calculate.Err)
	}

	if len(calculate.Uuid) == 0 {
		t.Fatalf("Uuid in responce must be not empty")
	}

	logrus.Info(calculate.Uuid)
}

func TestCalculator_Calculate_Invalid(t *testing.T) {
	calculator := Calculator{}

	_, err := calculator.Calculate(context.Background(), nil)
	if err == nil {
		t.Fatalf("Err must be 'request is nil', but err is void")
	}
}
