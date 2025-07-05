package port

import (
	"ExBot/internal/domain"
	"context"
)

// RequestRepository — CRUD-интерфейс для auth_requests
type RequestRepository interface {
	Save(ctx context.Context, r *domain.Request) error
	GetByID(ctx context.Context, id int64) (*domain.Request, error)
	ListByUser(ctx context.Context, userID int64) ([]*domain.Request, error)
	UpdateStatus(ctx context.Context, id int64, status string) error
}
