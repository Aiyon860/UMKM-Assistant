package handler

import (
	"encoding/json"
	"net/http"
	"todo-go/internal/model"
	"todo-go/internal/service"
)

type CaptionHandler struct {
	service service.CaptionService
}

func NewCaptionHandler(service service.CaptionService) *CaptionHandler {
	return &CaptionHandler{service: service}
}

func (h *CaptionHandler) GenerateCaption(w http.ResponseWriter, r *http.Request) {
	// Abaikan semua form input dan langsung kembalikan data mock
	caption, err := h.service.GenerateCaption(model.CaptionRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.CaptionResponse{
			Success: false,
			Message: "Gagal menghasilkan caption",
		})
		return
	}

	json.NewEncoder(w).Encode(model.CaptionResponse{
		Success: true,
		Caption: caption,
	})
}