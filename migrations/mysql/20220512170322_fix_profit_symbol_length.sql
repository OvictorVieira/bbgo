-- +up
ALTER TABLE profits
    CHANGE symbol symbol VARCHAR(20) NOT NULL;

-- +down
SELECT 1;
