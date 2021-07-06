package v1

import (
	"github.com/go-chi/chi/v5"
	"go_task/internal/app"
	"net/http"
)

type handler struct {
	app app.Application
}

func NewHandler(app app.Application, r chi.Router) http.Handler {
	return nil
}
