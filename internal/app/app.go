package app

import (
	"net/http"

	"github.com/imjowend/multimessenger/internal/domain"
	handler "github.com/imjowend/multimessenger/internal/infrastructure/delivery/http"
)

type App struct {
	handler *http.Handler
}

func NewApp(messageService *domain.MessageService) *App {
	handler := handler.NewHandler(messageService)
	return &App{handler: *http.Handler}
}

func (a *App) Run(addr string) error {
	http.HandleFunc("/webhook/whatsapp", a.handler.WhatsAppWebhook)
	return http.ListenAndServe(addr, nil)
}
