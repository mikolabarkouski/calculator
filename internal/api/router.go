package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/health", h.HealthCheck)
	r.Post("/calculate", h.calculate)

	r.Post("/package", h.addPackage)
	r.Delete("/package", h.deletePackage)
	r.Get("packages", h.getPackages)
	return r
}
