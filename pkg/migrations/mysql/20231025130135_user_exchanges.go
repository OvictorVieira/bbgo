package mysql

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upUserExchanges, downUserExchanges)

}

func upUserExchanges(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE `user_exchanges` (\n    `id`           BIGINT             PRIMARY KEY   NOT NULL AUTO_INCREMENT,\n    `user_id`      BIGINT                           NOT NULL,\n    `exchange_id`  BIGINT                           NOT NULL,\n    `api_key`      VARCHAR(255)       UNIQUE        NOT NULL,\n    `api_secret`   VARCHAR(255)       UNIQUE        NOT NULL,\n    `created_at`   TIMESTAMP          DEFAULT       (CURRENT_TIMESTAMP),\n    `updated_at`   TIMESTAMP          DEFAULT       (CURRENT_TIMESTAMP),\n    FOREIGN KEY (`user_id`) REFERENCES users(id),\n    FOREIGN KEY (`exchange_id`) REFERENCES exchanges(id)\n);")
	if err != nil {
		return err
	}

	return err
}

func downUserExchanges(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE `user_exchanges`;")
	if err != nil {
		return err
	}

	return err
}
