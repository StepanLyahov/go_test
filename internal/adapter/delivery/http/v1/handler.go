package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go_task/internal/adapter/delivery/http/httperr"
	"go_task/internal/app"
	"go_task/internal/app/command"
	"net/http"
)

type handler struct {
	app app.Application
}

func (h *handler) CalculatingId(w http.ResponseWriter, r *http.Request) {
	res, ok := unmarshallCalculatingUniqueIdRequest(w, r)
	if !ok {
		return
	}

	var mapped = command.CalculatingUniqueIdCommandRequest{
		Param1: res.Param1,
		Param2: res.Param2,
	}

	handle, err := h.app.Commands.CalculatingUniqueId.Handle(mapped)
	if err != nil {
		httperr.InternalServerError("Error Calculating Unique Id", err, w, r)
	}

	var result = CalculatingUniqueIdResponse{
		handle.Uuid,
	}

	buildResponse(result, w, http.StatusOK)
}

func NewHandler(app app.Application, r chi.Router) http.Handler {
	return HandlerFromMux(&handler{
		app: app,
	}, r)
}

func unmarshallCalculatingUniqueIdRequest(w http.ResponseWriter, r *http.Request) (CalculatingUniqueIdRequest, bool) {
	var rb CalculatingUniqueIdRequest
	if ok := decode(w, r, &rb); !ok {
		return CalculatingUniqueIdRequest{}, false
	}

	return rb, true
}

func decode(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := render.Decode(r, v); err != nil {
		httperr.BadRequest("bad-request", err, w, r)
		return false
	}

	return true
}

func buildResponse(v interface{}, w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
}
