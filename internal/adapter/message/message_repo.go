package message

import (
	"ExBot/internal/bot"
	"ExBot/internal/domain"
	"ExBot/internal/usecase"
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageRepo struct {
	bot        *tgbotapi.BotAPI
	MessageSvc *usecase.MessageService
}

func NewMessageRepo(bot *tgbotapi.BotAPI, messageSvc *usecase.MessageService) *MessageRepo {
	return &MessageRepo{bot: bot, MessageSvc: messageSvc}
}

func (m *MessageRepo) SendMessage(ctx context.Context, msg *domain.Message) error {
	tgMsg := tgbotapi.NewMessage(msg.ChatID, msg.Text)
	if msg.Keyboard != nil {
		tgMsg.ReplyMarkup = bot.ToTelegramMarkup(msg.Keyboard)
	}
	_, err := m.bot.Send(tgMsg)
	return err
}

func (m *MessageRepo) Listen(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := &domain.Message{
			UserID:    update.Message.From.ID,
			ChatID:    update.Message.Chat.ID,
			Text:      update.Message.Text,
			Timestamp: int64(update.Message.Date),
		}
		switch msg.Text {
		case "/start":
			_ = m.MessageSvc.HandleStart(ctx, msg)

		default:
			_ = m.MessageSvc.HandleEcho(ctx, msg)
		}
	}
}
