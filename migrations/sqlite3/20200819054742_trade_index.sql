-- +up
CREATE INDEX trades_symbol ON trades (user_exchanges_id, symbol);
CREATE INDEX trades_symbol_fee_currency ON trades (user_exchanges_id, symbol, fee_currency, traded_at);
CREATE INDEX trades_traded_at_symbol ON trades (user_exchanges_id, traded_at, symbol);

-- +down
DROP INDEX trades_symbol ON trades;
DROP INDEX trades_symbol_fee_currency ON trades;
DROP INDEX trades_traded_at_symbol ON trades;
