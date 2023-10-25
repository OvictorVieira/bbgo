-- +up
CREATE TABLE exchanges (
    id           INTEGER          PRIMARY KEY AUTOINCREMENT,
    name         TEXT             UNIQUE      NOT NULL,
    created_at   TIMESTAMP        DEFAULT     (datetime('now')),
    updated_at   TIMESTAMP        DEFAULT     (datetime('now'))
);

-- +down
DROP TABLE exchanges;
