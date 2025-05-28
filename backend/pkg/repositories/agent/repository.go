package agent_repo

import (
	"Arbitrax/pkg/services/bcrypt"
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, firstName, lastName, email, password string) (*Model, error)
	GetByUserUUID(ctx context.Context, uuid string) (*Model, error)
	GetAllByUserUUID(ctx context.Context, uuid string) (*Model, error)
}

type AgentRepository struct {
	db *pgxpool.Pool
}

func NewAgentRepo(db *pgxpool.Pool) *AgentRepository {
	return &AgentRepository{db: db}
}

func (r *AgentRepository) Create(ctx context.Context, firstName, lastName, email, password string) (*Model, error) {

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
