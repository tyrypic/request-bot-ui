package port

import (
	"ExBot/internal/domain"
	"context"
)

// UserRepository — CRUD-интерфейс для users
type UserRepository interface {
	Save(ctx context.Context, u *domain.User) error
	GetByTelegramID(ctx context.Context, tgID int64) (*domain.User, error)
	UpdateProfile(ctx context.Context, u *domain.User) error
	Approve(ctx context.Context, tgID int64) error
	SeedAdmin(ctx context.Context, adminID string) error
}
