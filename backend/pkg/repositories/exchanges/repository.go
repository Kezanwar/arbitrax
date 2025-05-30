package exchanges_repo

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*Model, error)
	GetByKey(ctx context.Context, key string) (*Model, error)
}

type ExchangeRepository struct {
	db *pgxpool.Pool
}

func NewExchangeRepo(db *pgxpool.Pool) *ExchangeRepository {
	return &ExchangeRepository{db: db}
}

func (r *ExchangeRepository) GetAll(ctx context.Context) ([]*Model, error) {
	var exchanges []*Model
	query := `SELECT * FROM exchanges`
	err := pgxscan.Select(ctx, r.db, &exchanges, query)
	if err != nil {
		return nil, fmt.Errorf("exchanges.GetAll: %w", err)
	}
	return exchanges, nil
}

func (r *ExchangeRepository) GetByKey(ctx context.Context, key string) (*Model, error) {
	var exchange Model
	query := `SELECT * FROM exchanges WHERE key=$1`
	err := pgxscan.Get(ctx, r.db, &exchange, query, key)
	if err != nil {
		return nil, fmt.Errorf("exchanges.GetByKey: %w", err)
	}
	return &exchange, nil
}
