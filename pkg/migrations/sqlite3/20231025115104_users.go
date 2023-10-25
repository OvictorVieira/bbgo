package sqlite3

import (
	"context"

	"github.com/c9s/rockhopper"
)

func init() {
	AddMigration(upUsers, downUsers)

}

func upUsers(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is applied.

	_, err = tx.ExecContext(ctx, "CREATE TABLE users (\n    id           INTEGER      PRIMARY KEY AUTOINCREMENT,\n    email        TEXT         UNIQUE      NOT NULL,\n    created_at   TIMESTAMP    DEFAULT     CURRENT_TIMESTAMP,\n    updated_at   TIMESTAMP    DEFAULT     CURRENT_TIMESTAMP\n);")
	if err != nil {
		return err
	}

	return err
}

func downUsers(ctx context.Context, tx rockhopper.SQLExecutor) (err error) {
	// This code is executed when the migration is rolled back.

	_, err = tx.ExecContext(ctx, "DROP TABLE users;")
	if err != nil {
		return err
	}

	return err
}
