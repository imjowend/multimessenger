package memory

import (
	"sync"

	"github.com/imjowend/multimessenger/internal/domain"
)

type messageRepository struct {
	messages map[string]*domain.Message
	mu       sync.Mutex
}

func NewMessageRepository() domain.MessageRepository {
	return &messageRepository{
		messages: make(map[string]*domain.Message),
	}
}

func (r *messageRepository) Save(message *domain.Message) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.messages[message.ID] = message
	return nil
}
