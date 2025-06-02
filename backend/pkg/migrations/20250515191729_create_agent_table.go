package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateAgentTable, downCreateAgentTable)
}

func upCreateAgentTable(ctx context.Context, tx *sql.Tx) error {
	createTable := `
      CREATE TABLE agents (
	   id SERIAL PRIMARY KEY,
	   uuid UUID DEFAULT uuid_generate_v4() UNIQUE,
	   user_uuid UUID NOT NULL,
	   name VARCHAR(100) NOT NULL,
	   avatar TEXT,
	   enabled BOOLEAN DEFAULT false,
	   capital_allocation DOUBLE PRECISION NOT NULL,
	   stop_loss DOUBLE PRECISION NOT NULL,
	   take_profit DOUBLE PRECISION NOT NULL,
	   exchanges TEXT[] NOT NULL,
	   strategies TEXT[] NOT NULL,
	   ai_orchestrated BOOLEAN DEFAULT false,
	   test_mode BOOLEAN DEFAULT false,
	   deleted BOOLEAN DEFAULT false,
	   created_at TIMESTAMP DEFAULT now(),
	   updated_at TIMESTAMP DEFAULT now()
      );
	  `

	if _, err := tx.Exec(createTable); err != nil {
		return fmt.Errorf("failed to create agents table: %w", err)
	}

	createIndex := `CREATE INDEX idx_agents_user_uuid ON agents(user_uuid)`
	if _, err := tx.Exec(createIndex); err != nil {
		return fmt.Errorf("failed to create index on user_uuid: %w", err)
	}

	return nil
}

func downCreateAgentTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE IF EXISTS agents`)
	return err
}
