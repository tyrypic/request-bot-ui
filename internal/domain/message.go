package domain

type Button struct {
	Text string // текст на кнопке
	Data string // payload (для Inline) или текст (для Reply)
}

// Keyboard описывает клавиатуру.
// Если Inline=true — будет InlineKeyboard, иначе обычная ReplyKeyboard.
type Keyboard struct {
	Inline  bool
	Buttons [][]Button
}

// Message — сущность для отправки.
type Message struct {
	ID        int64
	UserID    int64
	Text      string
	Timestamp int64
	ChatID    int64
	Keyboard  *Keyboard // nil, если без клавиатуры
}
