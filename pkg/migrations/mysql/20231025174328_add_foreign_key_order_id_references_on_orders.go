package mysql

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upAddForeignKeyOrderIdReferencesOnOrders, downAddForeignKeyOrderIdReferencesOnOrders)

}

func upAddForeignKeyOrderIdReferencesOnOrders(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "ALTER TABLE trades\n    ADD CONSTRAINT FK_Trades_Orders\n        FOREIGN KEY (order_id) REFERENCES orders(gid);")
	if err != nil {
		return err
	}

	return err
}

func downAddForeignKeyOrderIdReferencesOnOrders(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "ALTER TABLE trades DROP CONSTRAINT trades.order_id;")
	if err != nil {
		return err
	}

	return err
}
