package db

import (
	"ExBot/internal/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RequestRepo struct{ pool *pgxpool.Pool }

func NewRequestRepo(pool *pgxpool.Pool) *RequestRepo { return &RequestRepo{pool} }

// Save: INSERT или при conflict(user_id) полный UPDATE (UPSERT).
func (r *RequestRepo) Save(ctx context.Context, req *domain.Request) error {
	return r.pool.QueryRow(ctx, `
        INSERT INTO auth_requests
            (user_id,status,created_at)
        VALUES ($1,$2,NOW())
        ON CONFLICT (user_id) DO UPDATE
          SET status      = EXCLUDED.status,
              created_at  = EXCLUDED.created_at,
              resolved_at = NULL,
              resolved_by = NULL
        RETURNING id, created_at, status
    `, req.UserID, req.Status).
		Scan(&req.ID, &req.CreatedAt, &req.Status)
}

func (r *RequestRepo) UpdateStatus(ctx context.Context, id int64, status string) error {
	_, err := r.pool.Exec(ctx, `
        UPDATE auth_requests
           SET status=$2, resolved_at=NOW()
         WHERE id=$1
    `, id, status)
	return err
}

func (r *RequestRepo) GetByID(ctx context.Context, id int64) (*domain.Request, error) {
	req := &domain.Request{}
	err := r.pool.QueryRow(ctx, `
        SELECT id,user_id,status,created_at,resolved_at,resolved_by
          FROM auth_requests
         WHERE id=$1
    `, id).Scan(
		&req.ID, &req.UserID, &req.Status,
		&req.CreatedAt, &req.ResolvedAt, &req.ResolvedBy,
	)
	return req, err
}

func (r *RequestRepo) ListByUser(ctx context.Context, userID int64) ([]*domain.Request, error) {
	rows, err := r.pool.Query(ctx, `
        SELECT id,user_id,status,created_at,resolved_at,resolved_by
          FROM auth_requests
         WHERE user_id=$1
      ORDER BY created_at
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*domain.Request
	for rows.Next() {
		req := new(domain.Request)
		if err := rows.Scan(
			&req.ID, &req.UserID, &req.Status,
			&req.CreatedAt, &req.ResolvedAt, &req.ResolvedBy,
		); err != nil {
			return nil, err
		}
		out = append(out, req)
	}
	return out, rows.Err()
}
