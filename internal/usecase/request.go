package usecase

import (
	"ExBot/internal/domain"
	"ExBot/internal/port"
	"context"
	"time"
)

type RequestService struct {
	repo port.RequestRepository
}

func NewRequestService(r port.RequestRepository) *RequestService {
	return &RequestService{repo: r}
}

func (s *RequestService) Submit(ctx context.Context, userID int64) (*domain.Request, error) {
	r := &domain.Request{
		UserID:    userID,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	if err := s.repo.Save(ctx, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (s *RequestService) List(ctx context.Context, userID int64) ([]*domain.Request, error) {
	return s.repo.ListByUser(ctx, userID)
}

func (s *RequestService) Approve(ctx context.Context, id int64) error {
	return s.repo.UpdateStatus(ctx, id, "approved")
}

func (s *RequestService) Reject(ctx context.Context, id int64) error {
	return s.repo.UpdateStatus(ctx, id, "rejected")
}

func (s *RequestService) Get(ctx context.Context, id int64) (*domain.Request, error) {
	return s.repo.GetByID(ctx, id)
}
