-- +up
CREATE UNIQUE INDEX idx_kline_unique
    ON klines (`symbol`, `interval`, `start_time`);

-- +down
DROP INDEX `idx_kline_unique` ON `klines`;