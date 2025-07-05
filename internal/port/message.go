package port

import (
	"ExBot/internal/domain"
	"context"
)

type MessageRepository interface {
	SendMessage(ctx context.Context, msg *domain.Message) error
}
