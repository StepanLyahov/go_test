package command

type CalculatingUniqueIdHandler struct {
}

type CalculatingUniqueIdCommandRequest struct {
	Param1 string
	Param2 string
}

type CalculatingUniqueIdCommandResponse struct {
	Uuid string
}

func NewCalculatingUniqueIdHandler() CalculatingUniqueIdHandler {
	return CalculatingUniqueIdHandler{}
}

func (h *CalculatingUniqueIdHandler) Handle(request CalculatingUniqueIdCommandRequest) (CalculatingUniqueIdCommandResponse, error) {
	return CalculatingUniqueIdCommandResponse{}, nil
}
