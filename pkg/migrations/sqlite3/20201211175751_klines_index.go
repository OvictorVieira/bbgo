package sqlite3

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upKlinesIndex, downKlinesIndex)

}

func upKlinesIndex(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE INDEX IF NOT EXISTS `klines_end_time_symbol_interval` ON klines (`end_time`, `symbol`, `interval`);")
	if err != nil {
		return err
	}

	return err
}

func downKlinesIndex(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP INDEX IF EXISTS `klines_end_time_symbol_interval` ON `klines`;")
	if err != nil {
		return err
	}

	return err
}
