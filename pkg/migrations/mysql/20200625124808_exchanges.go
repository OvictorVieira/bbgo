package mysql

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upExchanges, downExchanges)

}

func upExchanges(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE `exchanges` (\n    `id`           BIGINT             PRIMARY KEY     NOT NULL AUTO_INCREMENT,\n    `name`         VARCHAR(255)       UNIQUE          NOT NULL,\n    `created_at`   TIMESTAMP          DEFAULT         (CURRENT_TIMESTAMP),\n    `updated_at`   TIMESTAMP          DEFAULT         (CURRENT_TIMESTAMP)\n);")
	if err != nil {
		return err
	}

	return err
}

func downExchanges(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE `exchanges`;")
	if err != nil {
		return err
	}

	return err
}
