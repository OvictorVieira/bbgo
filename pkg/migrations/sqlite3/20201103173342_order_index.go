package sqlite3

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upOrderIndex, downOrderIndex)

}

func upOrderIndex(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE INDEX orders_symbol ON orders (user_exchanges_id, symbol);")
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "CREATE UNIQUE INDEX orders_order_id ON orders (order_id, user_exchanges_id);")
	if err != nil {
		return err
	}

	return err
}

func downOrderIndex(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP INDEX orders_symbol ON orders;")
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DROP INDEX orders_order_id ON orders;")
	if err != nil {
		return err
	}

	return err
}
