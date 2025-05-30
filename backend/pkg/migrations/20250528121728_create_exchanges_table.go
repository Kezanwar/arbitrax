package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateExchangesTable, downCreateExchangesTable)
}

func upCreateExchangesTable(ctx context.Context, tx *sql.Tx) error {
	createTable := `
	CREATE TABLE exchanges (
		id SERIAL PRIMARY KEY,
		key VARCHAR(100) NOT NULL,
		label VARCHAR(100) NOT NULL
	);`

	if _, err := tx.ExecContext(ctx, createTable); err != nil {
		return fmt.Errorf("failed to create exchanges table: %w", err)
	}

	insertSeed := `
	INSERT INTO exchanges (key, label)
	VALUES 
		('ibkr', 'IBKR'),
		('binance', 'Binance'),
		('coinbase', 'Coinbase'),
		('kraken', 'Kraken'),
		('bybit', 'Bybit'),
		('bitfinex', 'Bitfinex');
	`

	if _, err := tx.ExecContext(ctx, insertSeed); err != nil {
		return fmt.Errorf("failed to seed exchanges: %w", err)
	}

	return nil
}

func downCreateExchangesTable(ctx context.Context, tx *sql.Tx) error {
	if _, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS exchanges`); err != nil {
		return fmt.Errorf("failed to drop exchanges table: %w", err)
	}
	return nil
}

/*
        ('Binance', 'https://logospng.org/wp-content/uploads/binance.png'),
		('Coinbase', 'https://www.liblogo.com/img-logo/co1496ca97-coinbase-logo-coinbase-promo-code-10-off-january-2022-the-wall-street-journal.png'),
		('Kraken', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQl8LFd933duAONmhOAtoGCgEq25AegXEprQg&s'),
		('Bybit', 'https://altcoinsbox.com/wp-content/uploads/2022/10/bybit-logo-white.jpg'),
		('Bitfinex', 'https://eu-images.contentstack.com/v3/assets/blt7dacf616844cf077/blte8f07cd9fa0abae2/679951ca333df67539f84793/bitfinex.png?width=1280&auto=webp&quality=95&format=jpg&disable=upscale');
*/
