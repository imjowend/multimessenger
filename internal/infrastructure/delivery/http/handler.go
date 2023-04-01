package http

import (
	"encoding/json"
	"net/http"

	"github.com/myapp/internal/domain"
)

type Handler struct {
	messageService *domain.MessageService
}

func NewHandler(messageService *domain.MessageService) *Handler {
	return &Handler{messageService: messageService}
}

func (h *Handler) WhatsAppWebhook(w http.ResponseWriter, r *http.Request) {
	var message domain.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.messageService.HandleMessage(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
