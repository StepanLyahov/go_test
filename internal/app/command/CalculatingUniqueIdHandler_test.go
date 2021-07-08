package command

import (
	"errors"
	"testing"
)

var str = "test"

type MockCalculatingWithoutErr struct {
}

func (m MockCalculatingWithoutErr) Execute(_, _ string) (string, error) {
	return str, nil
}

func TestCalculatingUniqueIdHandler_Execute_WithoutErr(t *testing.T) {
	cal := NewCalculatingUniqueIdHandler(MockCalculatingWithoutErr{})

	request := CalculatingUniqueIdCommandRequest{
		"param1",
		"param2",
	}

	handle, err := cal.Handle(request)
	if err != nil {
		t.Fatalf("Err must be nil, but: %v", err)
	}

	if str != handle.Uuid {
		t.Fatal("Both strings must be equal")
	}
}

type MockCalculatingWithErr struct {
}

func (m MockCalculatingWithErr) Execute(_, _ string) (string, error) {
	return "", errors.New("error")
}

func TestCalculatingUniqueIdHandler_Execute_WithErr(t *testing.T) {
	cal := NewCalculatingUniqueIdHandler(MockCalculatingWithErr{})

	request := CalculatingUniqueIdCommandRequest{
		"param1",
		"param2",
	}

	_, err := cal.Handle(request)
	if err == nil {
		t.Fatalf("Err must be exist, but err nil")
	}
}
