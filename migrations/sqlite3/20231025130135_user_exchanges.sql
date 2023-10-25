-- +up
CREATE TABLE user_exchanges (
    user_id      INTEGER          NOT NULL,
    exchange_id  INTEGER          NOT NULL,
    api_key      TEXT             UNIQUE NOT NULL,
    api_secret   TEXT             UNIQUE NOT NULL,
    created_at   TIMESTAMP        DEFAULT (datetime('now')),
    updated_at   TIMESTAMP        DEFAULT (datetime('now')),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (exchange_id) REFERENCES exchanges(id)
);

-- +down
DROP TABLE user_exchanges;
