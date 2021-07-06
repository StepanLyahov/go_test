package v1

import (
	"github.com/go-chi/chi/v5"
	"go_task/internal/app"
	"net/http"
)

type handler struct {
	app app.Application
}

func (h *handler) CalculatingId(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func NewHandler(app app.Application, r chi.Router) http.Handler {
	return HandlerFromMux(&handler{
		app: app,
	}, r)
}
