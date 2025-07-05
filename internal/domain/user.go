package domain

import "time"

// User — профиль пользователя
type User struct {
	ID         int64 // PK
	TelegramID int64 // Telegram ID
	Username   string
	FirstName  string
	LastName   string
	RealName   string // ЛК
	Email      string // ЛК
	Age        int    // ЛК
	City       string // ЛК
	IsAdmin    bool
	IsApproved bool
	CreatedAt  time.Time
	ApprovedAt *time.Time
}
