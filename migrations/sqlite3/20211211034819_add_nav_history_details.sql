-- +up
CREATE TABLE nav_history_details
(
    gid                  BIGINT             UNSIGNED        AUTO_INCREMENT          PRIMARY KEY,
    user_exchanges_id    BIGINT                                                     NOT NULL,
    subaccount           VARCHAR(30)                                                NOT NULL,
    time                 DATETIME(3)                                                NOT NULL,
    currency             VARCHAR(12)                                                NOT NULL,
    balance_in_usd       DECIMAL(32, 8)     UNSIGNED        DEFAULT 0.00000000      NOT NULL,
    balance_in_btc       DECIMAL(32, 8)     UNSIGNED        DEFAULT 0.00000000      NOT NULL,
    balance              DECIMAL(32, 8)     UNSIGNED        DEFAULT 0.00000000      NOT NULL,
    available            DECIMAL(32, 8)     UNSIGNED        DEFAULT 0.00000000      NOT NULL,
    locked               DECIMAL(32, 8)     UNSIGNED        DEFAULT 0.00000000      NOT NULL,

    FOREIGN KEY (`user_exchanges_id`) REFERENCES user_exchanges(id)
);

CREATE INDEX idx_nav_history_details ON nav_history_details (time, currency, user_exchanges_id);

-- +down
DROP TABLE nav_history_details;
