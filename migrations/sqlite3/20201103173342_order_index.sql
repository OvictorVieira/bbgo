-- +up
CREATE INDEX orders_symbol ON orders (user_exchanges_id, symbol);
CREATE UNIQUE INDEX orders_order_id ON orders (order_id, user_exchanges_id);

-- +down
DROP INDEX orders_symbol ON orders;
DROP INDEX orders_order_id ON orders;
