package db

import (
	"ExBot/internal/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct{ pool *pgxpool.Pool }

func NewUserRepo(pool *pgxpool.Pool) *UserRepo { return &UserRepo{pool} }

// Save: INSERT или UPDATE по telegram_id (ON CONFLICT).
func (r *UserRepo) Save(ctx context.Context, u *domain.User) error {
	return r.pool.QueryRow(ctx, `
        INSERT INTO users
            (telegram_id, username, first_name, last_name, created_at)
        VALUES ($1,$2,$3,$4,NOW())
        ON CONFLICT (telegram_id) DO UPDATE
          SET username   = EXCLUDED.username,
              first_name = EXCLUDED.first_name,
              last_name  = EXCLUDED.last_name
        RETURNING id, created_at, is_admin, approved_at
    `, u.TelegramID, u.Username, u.FirstName, u.LastName).
		Scan(&u.ID, &u.CreatedAt, &u.IsAdmin, &u.ApprovedAt)
}

func (r *UserRepo) GetByTelegramID(ctx context.Context, tgID int64) (*domain.User, error) {
	u := &domain.User{}
	err := r.pool.QueryRow(ctx, `
        SELECT
            id, telegram_id, username, first_name, last_name,
            real_name, email, age, city,
            is_admin, created_at, approved_at
        FROM users
        WHERE telegram_id=$1
    `, tgID).Scan(
		&u.ID, &u.TelegramID, &u.Username, &u.FirstName, &u.LastName,
		&u.RealName, &u.Email, &u.Age, &u.City,
		&u.IsAdmin, &u.CreatedAt, &u.ApprovedAt,
	)
	return u, err
}

// UpdateProfile: патч полей личного кабинета, остальные остаются.
func (r *UserRepo) UpdateProfile(ctx context.Context, u *domain.User) error {
	_, err := r.pool.Exec(ctx, `
        UPDATE users SET
          real_name = COALESCE(NULLIF($1,''), real_name),
          email     = COALESCE(NULLIF($2,''), email),
          age       = COALESCE(NULLIF($3::text,'')::int, age),
          city      = COALESCE(NULLIF($4,''), city)
        WHERE telegram_id=$5
    `,
		u.RealName, u.Email, u.Age, u.City, u.TelegramID,
	)
	return err
}

func (r *UserRepo) Approve(ctx context.Context, tgID int64) error {
	_, err := r.pool.Exec(ctx, `
        UPDATE users
           SET approved_at=NOW()
         WHERE telegram_id=$1
    `, tgID)
	return err
}

func (r *UserRepo) SeedAdmin(ctx context.Context, adminID string) error {
	if _, err := r.pool.Exec(ctx, `
            INSERT INTO users (telegram_id,is_admin,created_at)
            VALUES ($1,true,NOW())
            ON CONFLICT (telegram_id) DO UPDATE
              SET is_admin = true
        `, adminID); err != nil {
		return err
	}

	return nil
}
