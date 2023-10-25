-- +up
ALTER TABLE `klines`
    ADD COLUMN `quote_volume`           DECIMAL(32, 8) NOT NULL DEFAULT 0.0,
    ADD COLUMN `taker_buy_base_volume`  DECIMAL(32, 8) NOT NULL DEFAULT 0.0,
    ADD COLUMN `taker_buy_quote_volume` DECIMAL(32, 8) NOT NULL DEFAULT 0.0;

-- +down
ALTER TABLE `klines`
    DROP COLUMN `quote_volume`,
    DROP COLUMN `taker_buy_base_volume`,
    DROP COLUMN `taker_buy_quote_volume`;
