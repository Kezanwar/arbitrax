package agent_repo

import "time"

type Model struct {
	ID                int       `json:"-" db:"id"`
	UUID              string    `json:"uuid" db:"uuid"`
	UserUUID          string    `json:"-" db:"user_uuid"`
	Name              string    `json:"name" db:"name"`
	Avatar            string    `json:"avatar" db:"avatar"`
	Enabled           bool      `json:"enabled" db:"enabled"`
	CapitalAllocation float64   `json:"capital_allocation" db:"capital_allocation"`
	StopLoss          float64   `json:"stop_loss" db:"stop_loss"`
	TakeProfit        float64   `json:"take_profit" db:"take_profit"`
	Exchanges         []string  `json:"exchanges" db:"exchanges"`
	Strategies        []string  `json:"strategies" db:"strategies"`
	AiOrchestrated    bool      `json:"ai_orchestrated" db:"ai_orchestrated"`
	TestMode          bool      `json:"test_mode" db:"test_mode"`
	Deleted           bool      `json:"deleted" db:"deleted"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}
