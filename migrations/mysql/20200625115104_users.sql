-- +up
CREATE TABLE `users` (
    `id`           BIGINT           PRIMARY KEY     NOT NULL AUTO_INCREMENT,
    `email`        VARCHAR(255)     UNIQUE          NOT NULL,
    `created_at`   TIMESTAMP        DEFAULT         (CURRENT_TIMESTAMP),
    `updated_at`   TIMESTAMP        DEFAULT         (CURRENT_TIMESTAMP)
);

-- +down
DROP TABLE `users`;
