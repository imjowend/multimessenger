package main

import (
	"log"

	"github.com/myapp/internal/app"
	"github.com/myapp/internal/domain"
	"github.com/myapp/internal/infrastructure/repository/memory"
)

func main() {
	messageRepo := memory.NewMessageRepository()
	messageService := domain.NewMessageService(messageRepo)
	myApp := app.NewApp(messageService)

	addr := ":8080"
	log.Printf("Server listening on %s", addr)
	if err := myApp.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
