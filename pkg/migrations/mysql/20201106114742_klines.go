package mysql

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upKlines, downKlines)

}

func upKlines(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS `klines`\n(\n    `gid`                       BIGINT UNSIGNED                             NOT NULL AUTO_INCREMENT,\n    `user_exchanges_id`         BIGINT                                      NOT NULL,\n    `start_time`                DATETIME(3)                                 NOT NULL,\n    `end_time`                  DATETIME(3)                                 NOT NULL,\n    `interval`                  VARCHAR(3)                                  NOT NULL,\n    `symbol`                    VARCHAR(20)                                 NOT NULL,\n    `open`                      DECIMAL(20, 8)          UNSIGNED            NOT NULL,\n    `high`                      DECIMAL(20, 8)          UNSIGNED            NOT NULL,\n    `low`                       DECIMAL(20, 8)          UNSIGNED            NOT NULL,\n    `close`                     DECIMAL(20, 8)          UNSIGNED            NOT NULL DEFAULT '0.00000000',\n    `volume`                    DECIMAL(20, 8)          UNSIGNED            NOT NULL DEFAULT '0.00000000',\n    `closed`                    BOOL                                        NOT NULL DEFAULT TRUE,\n    `last_trade_id`             INT UNSIGNED                                NOT NULL DEFAULT 0,\n    `num_trades`                INT UNSIGNED                                NOT NULL DEFAULT 0,\n    `quote_volume`              DECIMAL(32, 8)                              NOT NULL DEFAULT '0.00000000',\n    `taker_buy_base_volume`     DECIMAL(32, 8)                              NOT NULL DEFAULT '0.00000000',\n    `taker_buy_quote_volume`    DECIMAL(32, 8)                              NOT NULL DEFAULT '0.00000000',\n    PRIMARY KEY (`gid`),\n    FOREIGN KEY (`user_exchanges_id`) REFERENCES user_exchanges(id)\n);")
	if err != nil {
		return err
	}

	return err
}

func downKlines(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE IF EXISTS `klines`;")
	if err != nil {
		return err
	}

	return err
}
