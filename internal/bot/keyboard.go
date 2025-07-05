package bot

import (
	"ExBot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewVerificationKeyboard() *domain.Keyboard {
	return &domain.Keyboard{
		Inline: true,
		Buttons: [][]domain.Button{
			{
				{Text: "Верификация", Data: "verify"},
			},
		},
	}
}

func ToTelegramMarkup(kb *domain.Keyboard) interface{} {
	if kb == nil {
		return nil
	}

	if kb.Inline {
		// Формируем InlineKeyboardMarkup
		rows := make([][]tgbotapi.InlineKeyboardButton, len(kb.Buttons))
		for i, row := range kb.Buttons {
			btns := make([]tgbotapi.InlineKeyboardButton, len(row))
			for j, btn := range row {
				btns[j] = tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.Data)
			}
			rows[i] = btns
		}
		return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
	}

	// Формируем обычную ReplyKeyboardMarkup
	rows := make([][]tgbotapi.KeyboardButton, len(kb.Buttons))
	for i, row := range kb.Buttons {
		btns := make([]tgbotapi.KeyboardButton, len(row))
		for j, btn := range row {
			btns[j] = tgbotapi.NewKeyboardButton(btn.Text)
		}
		rows[i] = btns
	}
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:        rows,
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
}
