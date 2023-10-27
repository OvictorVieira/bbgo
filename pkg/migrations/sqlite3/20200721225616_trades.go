package sqlite3

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upTrades, downTrades)

}

func upTrades(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS `trades`\n(\n    `gid`                   BIGINT UNSIGNED                     NOT NULL AUTO_INCREMENT,\n    `id`                    BIGINT UNSIGNED,\n    `order_id`              BIGINT UNSIGNED                     NOT NULL,\n    `user_exchanges_id`     BIGINT                              NOT NULL,\n    `symbol`                VARCHAR(20)                         NOT NULL,\n    `price`                 DECIMAL(16, 8)          UNSIGNED    NOT NULL,\n    `quantity`              DECIMAL(16, 8)          UNSIGNED    NOT NULL,\n    `quote_quantity`        DECIMAL(16, 8)          UNSIGNED    NOT NULL,\n    `fee`                   DECIMAL(16, 8)          UNSIGNED    NOT NULL,\n    `fee_currency`          BOOLEAN                             NOT NULL DEFAULT FALSE,\n    `side`                  VARCHAR(4)                          NOT NULL DEFAULT '',\n    `traded_at`             DATETIME(3)                         NOT NULL,\n    `is_margin`             BOOLEAN                             NOT NULL DEFAULT FALSE,\n    `is_isolated`           BOOLEAN                             NOT NULL DEFAULT FALSE,\n    `strategy`              VARCHAR(32)                         NULL,\n    `pnl`                   DECIMAL                             NULL,\n    PRIMARY KEY (`gid`),\n    UNIQUE KEY `id` (`user_exchanges_id`, `symbol`, `side`, `id`),\n    FOREIGN KEY (`user_exchanges_id`) REFERENCES user_exchanges(id)\n);")
	if err != nil {
		return err
	}

	return err
}

func downTrades(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE IF EXISTS `trades`;")
	if err != nil {
		return err
	}

	return err
}
