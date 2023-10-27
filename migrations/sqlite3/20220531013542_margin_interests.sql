-- +up
CREATE TABLE `margin_interests`
(
    `gid`                   BIGINT UNSIGNED          NOT NULL AUTO_INCREMENT,

    `user_exchanges_id`     BIGINT                   NOT NULL,

    `asset`                 VARCHAR(24)              NOT NULL DEFAULT '',

    `isolated_symbol`       VARCHAR(24)              NOT NULL DEFAULT '',

    `principle`             DECIMAL(16, 8) UNSIGNED  NOT NULL,

    `interest`              DECIMAL(20, 16) UNSIGNED NOT NULL,

    `interest_rate`         DECIMAL(20, 16) UNSIGNED NOT NULL,

    `time`                  DATETIME(3)              NOT NULL,

    PRIMARY KEY (`gid`),
    FOREIGN KEY (`user_exchanges_id`) REFERENCES user_exchanges(id)
);

-- +down
DROP TABLE IF EXISTS `margin_interests`;
