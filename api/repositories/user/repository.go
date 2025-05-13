package user_repo

import (
	"Arbitrax/db"
	"Arbitrax/services/bcrypt"
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, firstName, lastName, email, password string) (*Model, error)
	DoesEmailExist(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*Model, error)
	GetByUUID(ctx context.Context, uuid string) (*Model, error)
	FetchAll(ctx context.Context) ([]*Model, error)
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, firstName, lastName, email, password string) (*Model, error) {

	now := time.Now()

	hashPass, err := bcrypt.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("user.Create hashPw: %w", err)
	}

	query := `
		INSERT INTO users (first_name, last_name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
	`

	var user Model

	err = pgxscan.Get(ctx, r.db, &user, query, firstName, lastName, email, hashPass, now, now)

	if err != nil {
		return nil, fmt.Errorf("user.Create query: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`

	err := r.db.QueryRow(ctx, query, email).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("user.DoesEmailExist: %w", err)
	}

	return exists, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*Model, error) {
	var user Model

	query := `SELECT * FROM users WHERE email=$1`

	err := pgxscan.Get(ctx, r.db, &user, query, email)

	if err != nil {
		if db.IsNoRowsError(err) {
			return nil, fmt.Errorf("user.GetByEmail not found: %s", email)
		}

		return nil, fmt.Errorf("user.GetByEmail query: %w", err)
	}

	return &user, nil
}
func (r *UserRepository) GetByUUID(ctx context.Context, uuid string) (*Model, error) {
	var user Model
	query := `SELECT * FROM users WHERE uuid=$1`

	err := pgxscan.Get(ctx, r.db, &user, query, uuid)
	if err != nil {
		if db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("user.GetByUUID query: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) FetchAll(ctx context.Context) ([]*Model, error) {
	var users []*Model
	query := `SELECT * FROM users`

	err := pgxscan.Select(ctx, r.db, &users, query)
	if err != nil {
		return nil, fmt.Errorf("user.FetchAll query: %w", err)
	}

	return users, nil
}
