package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) addPackage(w http.ResponseWriter, r *http.Request) {
	var packageSize int
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return

	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &packageSize); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	h.app.AddPackage(packageSize)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deletePackage(w http.ResponseWriter, r *http.Request) {
	var id string
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return

	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &id); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	h.app.DeletePackage(id)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getPackages(w http.ResponseWriter, r *http.Request) {
	packages := h.app.GetPackages()
	responseBody, err := json.Marshal(packages)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
