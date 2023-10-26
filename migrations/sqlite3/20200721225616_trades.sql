-- +up
CREATE TABLE IF NOT EXISTS `trades`
(
    `gid`                   BIGINT UNSIGNED                     NOT NULL AUTO_INCREMENT,

    `id`                    BIGINT UNSIGNED,
    `order_id`              BIGINT UNSIGNED                     NOT NULL,
    `user_exchanges_id`     BIGINT                              NOT NULL,
    `symbol`                VARCHAR(20)                         NOT NULL,
    `price`                 DECIMAL(16, 8)          UNSIGNED    NOT NULL,
    `quantity`              DECIMAL(16, 8)          UNSIGNED    NOT NULL,
    `quote_quantity`        DECIMAL(16, 8)          UNSIGNED    NOT NULL,
    `fee`                   DECIMAL(16, 8)          UNSIGNED    NOT NULL,
    `fee_currency`          BOOLEAN                             NOT NULL DEFAULT FALSE,
    `side`                  VARCHAR(4)                          NOT NULL DEFAULT '',
    `traded_at`             DATETIME(3)                         NOT NULL,

    `is_margin`             BOOLEAN                             NOT NULL DEFAULT FALSE,
    `is_isolated`           BOOLEAN                             NOT NULL DEFAULT FALSE,

    `strategy`              VARCHAR(32)                         NULL,
    `pnl`                   DECIMAL                             NULL,

    PRIMARY KEY (`gid`),
    UNIQUE KEY `id` (`user_exchanges_id`, `symbol`, `side`, `id`),
    FOREIGN KEY (`user_exchanges_id`) REFERENCES user_exchanges(id)
);

-- +down
DROP TABLE IF EXISTS `trades`;
