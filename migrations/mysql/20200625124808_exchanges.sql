-- +up
CREATE TABLE `exchanges` (
    `id`           BIGINT             PRIMARY KEY     NOT NULL AUTO_INCREMENT,
    `name`         VARCHAR(255)       UNIQUE          NOT NULL,
    `created_at`   TIMESTAMP          DEFAULT         (CURRENT_TIMESTAMP),
    `updated_at`   TIMESTAMP          DEFAULT         (CURRENT_TIMESTAMP)
);

-- +down
DROP TABLE `exchanges`;

