package domain

type MessageService struct {
	repo MessageRepository
}

func NewMessageService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) HandleMessage(message *Message) error {
	return s.repo.Save(message)
}
