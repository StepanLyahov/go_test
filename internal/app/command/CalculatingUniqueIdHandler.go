package command

type CalculatorUuid interface {
	Execute(param1, param2 string) (string, error)
}

type CalculatingUniqueIdHandler struct {
	calculator CalculatorUuid
}

type CalculatingUniqueIdCommandRequest struct {
	Param1 string
	Param2 string
}

type CalculatingUniqueIdCommandResponse struct {
	Uuid string
}

func NewCalculatingUniqueIdHandler(calculator CalculatorUuid) CalculatingUniqueIdHandler {
	return CalculatingUniqueIdHandler{calculator}
}

func (h *CalculatingUniqueIdHandler) Handle(request CalculatingUniqueIdCommandRequest) (CalculatingUniqueIdCommandResponse, error) {

	execute, err := h.calculator.Execute(request.Param1, request.Param2)
	if err != nil {
		return CalculatingUniqueIdCommandResponse{}, err
	}

	return CalculatingUniqueIdCommandResponse{execute}, nil
}
