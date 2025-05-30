package trades_repo

import "time"

type Model struct {
	ID           int       `json:"-" db:"id"`
	PlatformUUID string    `json:"platform_uuid" db:"platform_uuid"`
	AgentUUID    string    `json:"agent_uuid" db:"agent_uuid"`
	Instrument   string    `json:"instrument" db:"instrument"`
	Price        int       `json:"price" db:"price"`
	Amount       int       `json:"amount" db:"amount"`
	Strategy     string    `json:"strategy" db:"strategy"`
	StopLoss     float64   `json:"stop_loss" db:"stop_loss"`
	TakeProfit   float64   `json:"take_profit" db:"take_profit"`
	Exchange     string    `json:"exchanges" db:"exchange"`
	TestMode     bool      `json:"test_mode" db:"test_mode"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
