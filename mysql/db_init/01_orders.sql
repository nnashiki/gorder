CREATE TABLE `orders`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
    `creator_id`    bigint unsigned DEFAULT NULL,
    `contractor_id` bigint unsigned DEFAULT NULL,
    `create_at`     datetime(3) DEFAULT NULL,
    `update_at`     datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY             `fk_orders_user_creator` (`creator_id`),
    KEY             `fk_orders_user_contractor` (`contractor_id`),
    CONSTRAINT `fk_orders_user_creator` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`),
    CONSTRAINT `fk_orders_user_contractor` FOREIGN KEY (`contractor_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
