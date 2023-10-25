-- +up
CREATE TABLE strategies (
    id              INTEGER         PRIMARY KEY AUTOINCREMENT,
    user_id         INTEGER         NOT NULL,
    strategy_name   TEXT            NOT NULL,
    strategy_config TEXT            NOT NULL,
    created_at      TIMESTAMP       DEFAULT     (datetime('now')),
    updated_at      TIMESTAMP       DEFAULT     (datetime('now')),

    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +down
DROP TABLE strategies;
