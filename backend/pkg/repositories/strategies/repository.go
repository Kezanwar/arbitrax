package strategy_repo

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*Model, error)
	GetByUUID(ctx context.Context, key string) (*Model, error)
	IsValid(ctx context.Context, key string) (bool, error)
	FilterInvalid(ctx context.Context, keys []string) ([]string, error)
}

type StrategyRepository struct {
	db *pgxpool.Pool
}

func NewStrategyRepo(db *pgxpool.Pool) *StrategyRepository {
	return &StrategyRepository{db: db}
}

func (r *StrategyRepository) GetAll(ctx context.Context) ([]*Model, error) {
	var strategies []*Model
	err := pgxscan.Select(ctx, r.db, &strategies, `SELECT * FROM strategies`)
	if err != nil {
		return nil, fmt.Errorf("strategy.GetAll query: %w", err)
	}
	return strategies, nil
}

func (r *StrategyRepository) GetByUUID(ctx context.Context, uuid string) (*Model, error) {
	var s Model
	err := pgxscan.Get(ctx, r.db, &s, `SELECT * FROM strategies WHERE uuid=$1`, uuid)
	if err != nil {
		return nil, fmt.Errorf("strategy.GetByUUID: %w", err)
	}
	return &s, nil
}

func (r *StrategyRepository) IsValid(ctx context.Context, uuid string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM strategies WHERE uuid=$1)`, uuid).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("strategy.IsValid query: %w", err)
	}
	return exists, nil
}

func (r *StrategyRepository) FilterInvalid(ctx context.Context, keys []string) ([]string, error) {
	var valid []string
	for _, k := range keys {
		ok, err := r.IsValid(ctx, k)
		if err != nil {
			return nil, err
		}
		if ok {
			valid = append(valid, k)
		}
	}
	return valid, nil
}
