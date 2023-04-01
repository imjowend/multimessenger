package domain

type Message struct {
	ID        string
	Sender    string
	Recipient string
	Content   string
}

type MessageRepository interface {
	Save(message *Message) error
}
