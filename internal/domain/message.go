package domain

type Message struct {
	ID        int64
	UserID    int64
	Text      string
	Timestamp int64
	ChatID    int64
}
