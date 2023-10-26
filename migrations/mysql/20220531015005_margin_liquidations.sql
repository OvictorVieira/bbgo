-- +up
CREATE TABLE `margin_liquidations`
(
    `gid`                       BIGINT UNSIGNED         NOT NULL AUTO_INCREMENT,

    `user_exchanges_id`         BIGINT                  NOT NULL,

    `symbol`                    VARCHAR(24)             NOT NULL DEFAULT '',

    `order_id`                  BIGINT UNSIGNED         NOT NULL,

    `is_isolated`               BOOL                    NOT NULL DEFAULT false,

    `average_price`             DECIMAL(16, 8) UNSIGNED NOT NULL,

    `price`                     DECIMAL(16, 8) UNSIGNED NOT NULL,

    `quantity`                  DECIMAL(16, 8) UNSIGNED NOT NULL,

    `executed_quantity`         DECIMAL(16, 8) UNSIGNED NOT NULL,

    `side`                      VARCHAR(5)              NOT NULL DEFAULT '',

    `time_in_force`             VARCHAR(5)              NOT NULL DEFAULT '',

    `time`                      DATETIME(3)             NOT NULL,

    PRIMARY KEY (`gid`),
    UNIQUE KEY (`order_id`, `user_exchanges_id`),
    FOREIGN KEY (`user_exchanges_id`) REFERENCES user_exchanges(id)
);

-- +down
DROP TABLE IF EXISTS `margin_liquidations`;
