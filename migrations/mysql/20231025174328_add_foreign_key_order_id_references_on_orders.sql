-- +up
ALTER TABLE trades
    ADD CONSTRAINT FK_Trades_Orders
        FOREIGN KEY (order_id) REFERENCES orders(gid);

-- +down
ALTER TABLE trades DROP CONSTRAINT trades.order_id;
