-- +up
CREATE TABLE `user_exchanges` (
    `id`           BIGINT             PRIMARY KEY   NOT NULL AUTO_INCREMENT,
    `user_id`      BIGINT                           NOT NULL,
    `exchange_id`  BIGINT                           NOT NULL,
    `api_key`      VARCHAR(255)       UNIQUE        NOT NULL,
    `api_secret`   VARCHAR(255)       UNIQUE        NOT NULL,
    `created_at`   TIMESTAMP          DEFAULT       (CURRENT_TIMESTAMP),
    `updated_at`   TIMESTAMP          DEFAULT       (CURRENT_TIMESTAMP),
    FOREIGN KEY (`user_id`) REFERENCES users(id),
    FOREIGN KEY (`exchange_id`) REFERENCES exchanges(id)

);

-- +down
DROP TABLE `user_exchanges`;

