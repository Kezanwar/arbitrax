package agent_repo

import (
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Create(
		ctx context.Context,
		userUUID, name, avatar string,
		enabled bool,
		capitalAllocation, stopLoss, takeProfit float64,
		exchanges, strategies []string,
		aiOrchestrated, testMode bool,
	) (*Model, error)
	Update(
		ctx context.Context,
		uuid, name, avatar string,
		enabled bool,
		capitalAllocation, stopLoss, takeProfit float64,
		exchanges, strategies []string,
		aiOrchestrated, testMode bool,
	) (*Model, error)
	GetByUUID(ctx context.Context, uuid string) (*Model, error)
	GetAllByUserUUID(ctx context.Context, userUUID string) ([]*Model, error)
}

type AgentRepository struct {
	db *pgxpool.Pool
}

func NewAgentRepo(db *pgxpool.Pool) *AgentRepository {
	return &AgentRepository{db: db}
}

func (r *AgentRepository) Create(
	ctx context.Context,
	userUUID, name, avatar string,
	enabled bool,
	capitalAllocation, stopLoss, takeProfit float64,
	exchanges, strategies []string,
	aiOrchestrated, testMode bool,
) (*Model, error) {

	now := time.Now()

	query := `
		INSERT INTO agents (user_uuid, name, avatar, enabled, capital_allocation, stop_loss, take_profit, exchanges, strategies, ai_orchestrated, test_mode, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING *
	`

	var agent Model

	err := pgxscan.Get(ctx, r.db, &agent, query, userUUID, name, avatar, enabled, capitalAllocation, stopLoss, takeProfit, exchanges, strategies, aiOrchestrated, testMode, now, now)

	if err != nil {
		return nil, fmt.Errorf("agent.Create query: %w", err)
	}

	return &agent, nil
}

func (r *AgentRepository) GetByUUID(ctx context.Context, uuid string) (*Model, error) {
	query := `
		SELECT * FROM agents
		WHERE uuid = $1 AND deleted = false
	`

	var agent Model

	err := pgxscan.Get(ctx, r.db, &agent, query, uuid)
	if err != nil {
		return nil, fmt.Errorf("agent.GetByUUID query: %w", err)
	}

	return &agent, nil
}

func (r *AgentRepository) Update(
	ctx context.Context,
	uuid, name, avatar string,
	enabled bool,
	capitalAllocation, stopLoss, takeProfit float64,
	exchanges, strategies []string,
	aiOrchestrated, testMode bool,
) (*Model, error) {

	now := time.Now()

	query := `
		UPDATE agents 
		SET name = $2, avatar = $3, enabled = $4, capital_allocation = $5, 
		    stop_loss = $6, take_profit = $7, exchanges = $8, strategies = $9, 
		    ai_orchestrated = $10, test_mode = $11, updated_at = $12
		WHERE uuid = $1
		RETURNING *
	`

	var agent Model

	err := pgxscan.Get(ctx, r.db, &agent, query, uuid, name, avatar, enabled, capitalAllocation, stopLoss, takeProfit, exchanges, strategies, aiOrchestrated, testMode, now)

	if err != nil {
		return nil, fmt.Errorf("agent.Update query: %w", err)
	}

	return &agent, nil
}

func (r *AgentRepository) GetAllByUserUUID(ctx context.Context, userUUID string) ([]*Model, error) {
	query := `
		SELECT * FROM agents
		WHERE user_uuid = $1 AND deleted = false
		ORDER BY created_at DESC
	`

	agents := make([]*Model, 0)

	err := pgxscan.Select(ctx, r.db, &agents, query, userUUID)
	if err != nil {
		return nil, fmt.Errorf("agent.GetByUserUUID query: %w", err)
	}

	return agents, nil
}
