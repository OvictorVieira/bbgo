-- +up
CREATE TABLE `withdraws`
(
    `gid`                       BIGINT      UNSIGNED    NOT NULL AUTO_INCREMENT,
    `user_exchanges_id`         BIGINT                  NOT NULL,

    -- asset is the asset name (currency)
    `asset`                     VARCHAR(10)             NOT NULL,

    `address`                   VARCHAR(128)            NOT NULL,
    `network`                   VARCHAR(32)             NOT NULL DEFAULT '',

    `amount`                    DECIMAL(16, 8)          NOT NULL,
    `txn_id`                    VARCHAR(256)            NOT NULL,
    `txn_fee`                   DECIMAL(16, 8)          NOT NULL DEFAULT 0,
    `txn_fee_currency`          VARCHAR(32)             NOT NULL DEFAULT '',
    `time`                      DATETIME(3)             NOT NULL,

    PRIMARY KEY (`gid`),
    UNIQUE KEY `txn_id` (`user_exchanges_id`, `txn_id`),
    FOREIGN KEY (`user_exchanges_id`) REFERENCES user_exchanges(id)
);

-- +down
DROP TABLE IF EXISTS `withdraws`;
