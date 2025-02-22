package api

import (
	"net/http"

	"github.com/mikolabarkouski/calculator/internal/app"
)

type Handler struct {
	app *app.App
}

func NewHandler(a *app.App) *Handler {
	return &Handler{app: a}
}

func (h *Handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.HealthCheck)
	return mux
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
