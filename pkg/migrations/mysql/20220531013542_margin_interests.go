package mysql

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upMarginInterests, downMarginInterests)

}

func upMarginInterests(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE `margin_interests`\n(\n    `gid`             BIGINT UNSIGNED          NOT NULL AUTO_INCREMENT,\n    `exchange_id`     BIGINT                   NOT NULL DEFAULT '',\n    `asset`           VARCHAR(24)              NOT NULL DEFAULT '',\n    `isolated_symbol` VARCHAR(24)              NOT NULL DEFAULT '',\n    `principle`       DECIMAL(16, 8) UNSIGNED  NOT NULL,\n    `interest`        DECIMAL(20, 16) UNSIGNED NOT NULL,\n    `interest_rate`   DECIMAL(20, 16) UNSIGNED NOT NULL,\n    `time`            DATETIME(3)              NOT NULL,\n    PRIMARY KEY (`gid`)\n);")
	if err != nil {
		return err
	}

	return err
}

func downMarginInterests(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE IF EXISTS `margin_interests`;")
	if err != nil {
		return err
	}

	return err
}
