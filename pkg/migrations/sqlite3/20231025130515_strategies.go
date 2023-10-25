package sqlite3

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upStrategies, downStrategies)

}

func upStrategies(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE strategies (\n    id              INTEGER         PRIMARY KEY AUTOINCREMENT,\n    user_id         INTEGER         NOT NULL,\n    strategy_name   TEXT            NOT NULL,\n    strategy_config TEXT            NOT NULL,\n    created_at      TIMESTAMP       DEFAULT     (datetime('now')),\n    updated_at      TIMESTAMP       DEFAULT     (datetime('now')),\n    FOREIGN KEY (user_id) REFERENCES users(id)\n);")
	if err != nil {
		return err
	}

	return err
}

func downStrategies(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE strategies;")
	if err != nil {
		return err
	}

	return err
}
