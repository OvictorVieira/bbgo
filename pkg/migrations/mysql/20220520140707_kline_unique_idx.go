package mysql

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upKlineUniqueIdx, downKlineUniqueIdx)

}

func upKlineUniqueIdx(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE UNIQUE INDEX idx_kline_unique\n    ON klines (`symbol`, `interval`, `start_time`);")
	if err != nil {
		return err
	}

	return err
}

func downKlineUniqueIdx(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP INDEX `idx_kline_unique` ON `klines`;")
	if err != nil {
		return err
	}

	return err
}
