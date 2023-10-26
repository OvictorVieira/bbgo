-- +up
CREATE TABLE `strategies` (
    `id`              BIGINT           PRIMARY KEY      NOT NULL AUTO_INCREMENT,
    `user_id`         BIGINT                            NOT NULL,
    `strategy_name`   VARCHAR(255)                      NOT NULL,
    `strategy_config` JSON                              NOT NULL,
    `created_at`      TIMESTAMP        DEFAULT          (CURRENT_TIMESTAMP),
    `updated_at`      TIMESTAMP        DEFAULT          (CURRENT_TIMESTAMP),

    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);

-- +down
DROP TABLE `strategies`;

