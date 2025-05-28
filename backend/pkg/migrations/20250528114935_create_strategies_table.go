package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateStrategiesTable, downCreateStrategiesTable)
}

func upCreateStrategiesTable(ctx context.Context, tx *sql.Tx) error {
	createTable := `
	CREATE TABLE strategies (
		id SERIAL PRIMARY KEY,
		uuid UUID DEFAULT uuid_generate_v4() UNIQUE,
		key VARCHAR(100) NOT NULL,
		label VARCHAR(100) NOT NULL,
		description TEXT NOT NULL
	);`

	if _, err := tx.ExecContext(ctx, createTable); err != nil {
		return fmt.Errorf("failed to create strategies table: %w", err)
	}

	createIndex := `CREATE INDEX idx_strategies_key ON strategies(uuid);`
	if _, err := tx.ExecContext(ctx, createIndex); err != nil {
		return fmt.Errorf("failed to create index on strategies.key: %w", err)
	}

	// Seed data
	type seed struct {
		key, label, description string
	}
	seeds := []seed{
		{"olivia", "Olivia", "A conservative trend-following strategy focused on long-term growth."},
		{"jack", "Jack", "An aggressive mean-reversion strategy designed for high volatility."},
		{"emma", "Emma", "A balanced breakout strategy ideal for sideway markets."},
		{"liam", "Liam", "A scalping strategy that excels in high-frequency environments."},
		{"sophia", "Sophia", "A momentum strategy tailored for bullish trends."},
	}

	for _, s := range seeds {

		_, err := tx.ExecContext(ctx, `
			INSERT INTO strategies (key, label, description)
			VALUES ($1, $2, $3)`,
			s.key, s.label, s.description,
		)
		if err != nil {
			return fmt.Errorf("failed to seed strategy '%s': %w", s.label, err)
		}
	}

	return nil
}

func downCreateStrategiesTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS strategies`)
	return err
}
