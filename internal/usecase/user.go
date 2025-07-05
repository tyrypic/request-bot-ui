package usecase

import (
	"ExBot/internal/domain"
	"ExBot/internal/port"
	"context"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(r port.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) SeedAdmin(ctx context.Context, adminID int64) error {
	return s.repo.SeedAdmin(ctx, adminID)
}

func (s *UserService) Register(ctx context.Context, u *domain.User) error {
	return s.repo.Save(ctx, u)
}

func (s *UserService) UpdateProfile(ctx context.Context, u *domain.User) error {
	return s.repo.UpdateProfile(ctx, u)
}

func (s *UserService) Approve(ctx context.Context, tgID int64) error {
	return s.repo.Approve(ctx, tgID)
}

func (s *UserService) GetByTelegramID(ctx context.Context, tgID int64) (*domain.User, error) {
	return s.repo.GetByTelegramID(ctx, tgID)
}
