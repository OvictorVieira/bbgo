-- +up
ALTER TABLE `klines`
MODIFY COLUMN `volume` decimal(20,8) unsigned NOT NULL DEFAULT '0.00000000';

-- +down
ALTER TABLE `klines`
MODIFY COLUMN `volume` decimal(16,8) unsigned NOT NULL DEFAULT '0.00000000';
