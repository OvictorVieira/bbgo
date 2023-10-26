-- +up
ALTER TABLE trades
    CHANGE fee fee DECIMAL(16, 8) NOT NULL;

ALTER TABLE profits
    CHANGE fee fee DECIMAL(16, 8) NOT NULL;

ALTER TABLE profits
    CHANGE fee_in_usd fee_in_usd DECIMAL(16, 8);

-- +down
SELECT 1;
