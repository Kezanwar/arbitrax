package user_repo

import (
	"Arbitrax/db"
	"Arbitrax/services/bcrypt"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, firstName, lastName, email, password string) (*Model, error)
	DoesEmailExist(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*Model, error)
	GetByUUID(ctx context.Context, uuid string) (*Model, error)
	FetchAll(ctx context.Context) ([]*Model, error)
}

type PgxUserRepository struct {
	db *pgxpool.Pool
}

func NewPgxUserRepo(db *pgxpool.Pool) *PgxUserRepository {
	return &PgxUserRepository{db: db}
}

func (r *PgxUserRepository) Create(ctx context.Context, firstName, lastName, email, password string) (*Model, error) {
	now := time.Now()
	hashPass, err := bcrypt.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("user.Create hashPw: %w", err)
	}

	query := `
		INSERT INTO users (first_name, last_name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, uuid, first_name, last_name, email, password, created_at, updated_at;
	`

	row := r.db.QueryRow(ctx, query, firstName, lastName, email, hashPass, now, now)
	user := &Model{}
	if err := ScanIntoUser(row, user); err != nil {
		return nil, fmt.Errorf("user.Create scan: %w", err)
	}
	return user, nil
}

func (r *PgxUserRepository) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	query := `SELECT email FROM users WHERE email=$1`
	row := r.db.QueryRow(ctx, query, email)

	var found string
	err := row.Scan(&found)
	if err != nil {
		if db.IsNoRowsError(err) {
			return false, nil
		}
		return false, fmt.Errorf("user.DoesEmailExist scan: %w", err)
	}
	return true, nil
}

func (r *PgxUserRepository) GetByEmail(ctx context.Context, email string) (*Model, error) {
	query := `
		SELECT id, uuid, first_name, last_name, email, password, created_at, updated_at
		FROM users WHERE email=$1
	`
	row := r.db.QueryRow(ctx, query, email)
	user := &Model{}
	if err := ScanIntoUser(row, user); err != nil {
		if db.IsNoRowsError(err) {
			return nil, fmt.Errorf("user.GetByEmail not found: %s", email)
		}
		return nil, fmt.Errorf("user.GetByEmail scan: %w", err)
	}
	return user, nil
}

func (r *PgxUserRepository) GetByUUID(ctx context.Context, uuid string) (*Model, error) {
	query := `
		SELECT id, uuid, first_name, last_name, email, password, created_at, updated_at
		FROM users WHERE uuid=$1
	`
	row := r.db.QueryRow(ctx, query, uuid)
	user := &Model{}
	if err := ScanIntoUser(row, user); err != nil {
		if db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("user.GetByUUID scan: %w", err)
	}
	return user, nil
}

func (r *PgxUserRepository) FetchAll(ctx context.Context) ([]*Model, error) {
	query := `
		SELECT id, uuid, first_name, last_name, email, password, created_at, updated_at
		FROM users
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("user.FetchAll query: %w", err)
	}
	defer rows.Close()

	var users []*Model
	for rows.Next() {
		user := &Model{}
		if err := ScanIntoUser(rows, user); err != nil {
			return nil, fmt.Errorf("user.FetchAll scan: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("user.FetchAll rows: %w", err)
	}
	return users, nil
}
