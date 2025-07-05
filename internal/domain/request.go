package domain

import "time"

// Request — заявка на вступление
type Request struct {
	ID         int64  // PK
	UserID     int64  // FK → users.id
	Status     string // pending, approved, rejected
	CreatedAt  time.Time
	ResolvedAt *time.Time
	ResolvedBy *int64 // пользователь-админ
}
