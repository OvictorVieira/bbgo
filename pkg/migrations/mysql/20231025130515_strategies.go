package mysql

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upStrategies, downStrategies)

}

func upStrategies(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE `strategies` (\n    `id`              BIGINT           PRIMARY KEY      NOT NULL AUTO_INCREMENT,\n    `user_id`         BIGINT                            NOT NULL,\n    `strategy_name`   VARCHAR(255)                      NOT NULL,\n    `strategy_config` JSON                              NOT NULL,\n    `created_at`      TIMESTAMP        DEFAULT          (CURRENT_TIMESTAMP),\n    `updated_at`      TIMESTAMP        DEFAULT          (CURRENT_TIMESTAMP),\n    FOREIGN KEY (`user_id`) REFERENCES users(`id`)\n);")
	if err != nil {
		return err
	}

	return err
}

func downStrategies(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE `strategies`;")
	if err != nil {
		return err
	}

	return err
}
