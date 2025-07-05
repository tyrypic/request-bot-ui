package usecase

import (
	"ExBot/internal/domain"
	"ExBot/internal/port"
	"context"
)

type MessageService struct {
	userSvc port.UserRepository
	sender  port.MessageRepository
}

func NewMessageService(userSvc port.UserRepository, sender port.MessageRepository) *MessageService {
	return &MessageService{userSvc: userSvc, sender: sender}
}

func (s *MessageService) HandleStart(ctx context.Context, msg *domain.Message) error {
	user, err := s.userSvc.GetByTelegramID(ctx, msg.UserID)
	if err != nil {
		msg.Text = "Ошибка при проверке пользователя"
	} else if user == nil {
		msg.Text = "Вы не администратор."
	} else if user.IsAdmin {
		msg.Text = "Вы администратор!"
	} else {
		msg.Text = "Вы не администратор."
	}
	return s.sender.SendMessage(ctx, msg)
}

func (s *MessageService) HandleEcho(ctx context.Context, msg *domain.Message) error {
	msg.Text = "Эхо: " + msg.Text
	return s.sender.SendMessage(ctx, msg)
}
