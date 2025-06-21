package handler

import (
	"encoding/json"
	"net/http"
	"todo-go/internal/model"
	"todo-go/internal/service"
)

type ChatbotHandler struct {
	service *service.ChatbotService
}

func NewChatbotHandler(service *service.ChatbotService) *ChatbotHandler {
	return &ChatbotHandler{service: service}
}

func (h *ChatbotHandler) Chat(w http.ResponseWriter, r *http.Request) {
	var req model.ChatRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Message == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp := h.service.Chat(req.Message)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}