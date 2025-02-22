package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) calculate(w http.ResponseWriter, r *http.Request) {
	var orderSizeRequest int
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &orderSizeRequest); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	packageSizes := h.app.GetPackages()

	result, err := h.app.CalculatePacksNeeded(orderSizeRequest, packageSizes)
	if err != nil {
		http.Error(w, "error happened", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &orderSizeRequest)
	if err != nil {
		http.Error(w, "error happened", http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
