-- +up
CREATE INDEX `klines_end_time_symbol_interval` ON klines (`end_time`, `symbol`, `interval`);

-- +down
DROP INDEX `klines_end_time_symbol_interval` ON `klines`;
