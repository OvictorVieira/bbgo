-- +up
ALTER TABLE `trades` ADD COLUMN `is_futures` BOOLEAN NOT NULL DEFAULT FALSE;
ALTER TABLE `orders` ADD COLUMN `is_futures` BOOLEAN NOT NULL DEFAULT FALSE;

-- +down
ALTER TABLE `trades` DROP COLUMN `is_futures`;
ALTER TABLE `orders` DROP COLUMN `is_futures`;
