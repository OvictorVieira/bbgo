-- +up
CREATE TABLE nav_history_details
(
    gid            bigint             unsigned auto_increment PRIMARY KEY,
    exchange_id    bigint                                     NOT NULL,
    subaccount     VARCHAR(30)                                NOT NULL,
    time           DATETIME(3)                                NOT NULL,
    currency       VARCHAR(12)                                NOT NULL,
    balance_in_usd DECIMAL(32, 8) UNSIGNED DEFAULT 0.00000000 NOT NULL,
    balance_in_btc DECIMAL(32, 8) UNSIGNED DEFAULT 0.00000000 NOT NULL,
    balance        DECIMAL(32, 8) UNSIGNED DEFAULT 0.00000000 NOT NULL,
    available      DECIMAL(32, 8) UNSIGNED DEFAULT 0.00000000 NOT NULL,
    locked         DECIMAL(32, 8) UNSIGNED DEFAULT 0.00000000 NOT NULL
);

CREATE INDEX idx_nav_history_details
    on nav_history_details (time, currency, exchange_id);

-- +down
DROP TABLE nav_history_details;
