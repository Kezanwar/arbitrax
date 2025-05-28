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
		label VARCHAR(100) NOT NULL,
		avatar_url TEXT NOT NULL,
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
		label, description string
	}
	seeds := []seed{
		{"Olivia", "A conservative trend-following strategy focused on long-term growth."},
		{"Jack", "An aggressive mean-reversion strategy designed for high volatility."},
		{"Emma", "A balanced breakout strategy ideal for sideway markets."},
		{"Liam", "A scalping strategy that excels in high-frequency environments."},
		{"Sophia", "A momentum strategy tailored for bullish trends."},
	}

	for _, s := range seeds {
		avatarURL := fmt.Sprintf("https://robohash.org/%s.png?size=200x200&set=set1", s.label)
		_, err := tx.ExecContext(ctx, `
			INSERT INTO strategies (label, avatar_url, description)
			VALUES ($1, $2, $3)`,
			s.label, avatarURL, s.description,
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
