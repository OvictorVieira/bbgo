-- +up
CREATE TABLE IF NOT EXISTS `klines`
(
    `gid`                       BIGINT UNSIGNED                             NOT NULL AUTO_INCREMENT,
    `user_exchanges_id`         BIGINT                                      NOT NULL,
    `start_time`                DATETIME(3)                                 NOT NULL,
    `end_time`                  DATETIME(3)                                 NOT NULL,
    `interval`                  VARCHAR(3)                                  NOT NULL,
    `symbol`                    VARCHAR(20)                                 NOT NULL,
    `open`                      DECIMAL(20, 8)          UNSIGNED            NOT NULL,
    `high`                      DECIMAL(20, 8)          UNSIGNED            NOT NULL,
    `low`                       DECIMAL(20, 8)          UNSIGNED            NOT NULL,
    `close`                     DECIMAL(20, 8)          UNSIGNED            NOT NULL DEFAULT '0.00000000',
    `volume`                    DECIMAL(20, 8)          UNSIGNED            NOT NULL DEFAULT '0.00000000',
    `closed`                    BOOL                                        NOT NULL DEFAULT TRUE,
    `last_trade_id`             INT UNSIGNED                                NOT NULL DEFAULT 0,
    `num_trades`                INT UNSIGNED                                NOT NULL DEFAULT 0,
    `quote_volume`              DECIMAL(32, 8)                              NOT NULL DEFAULT '0.00000000',
    `taker_buy_base_volume`     DECIMAL(32, 8)                              NOT NULL DEFAULT '0.00000000',
    `taker_buy_quote_volume`    DECIMAL(32, 8)                              NOT NULL DEFAULT '0.00000000',

    PRIMARY KEY (`gid`),
    FOREIGN KEY (`user_exchanges_id`) REFERENCES user_exchanges(id)
);

-- +down
DROP TABLE IF EXISTS `klines`;
