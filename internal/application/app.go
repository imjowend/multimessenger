package application

import (
	"net/http"

	"github.com/myapp/internal/domain"
	"github.com/myapp/internal/infrastructure/delivery/http"
)

type App struct {
	handler *http.Handler
}

func NewApp(messageService *domain.MessageService) *App {
	handler := http.NewHandler(messageService)
	return &App{handler: handler}
}

func (a *App) Run(addr string) error {
	http.HandleFunc("/webhook/whatsapp", a.handler.WhatsAppWebhook)
	return http.ListenAndServe(addr, nil)
}
