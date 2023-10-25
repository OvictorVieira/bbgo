package sqlite3

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upUserExchanges, downUserExchanges)

}

func upUserExchanges(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE user_exchanges (\n    user_id      INTEGER          NOT NULL,\n    exchange_id  INTEGER          NOT NULL,\n    api_key      TEXT             UNIQUE NOT NULL,\n    api_secret   TEXT             UNIQUE NOT NULL,\n    created_at   TIMESTAMP        DEFAULT (datetime('now')),\n    updated_at   TIMESTAMP        DEFAULT (datetime('now')),\n    FOREIGN KEY (user_id) REFERENCES users(id),\n    FOREIGN KEY (exchange_id) REFERENCES exchanges(id)\n);")
	if err != nil {
		return err
	}

	return err
}

func downUserExchanges(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE user_exchanges;")
	if err != nil {
		return err
	}

	return err
}
