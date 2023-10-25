-- +up
CREATE TABLE users (
    id           INTEGER      PRIMARY KEY AUTOINCREMENT,
    email        TEXT         UNIQUE      NOT NULL,
    created_at   TIMESTAMP    DEFAULT     CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    DEFAULT     CURRENT_TIMESTAMP
);

-- +down
DROP TABLE users;
