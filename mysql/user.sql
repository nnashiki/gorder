CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`       varchar(255) NOT NULL,
    `email`      varchar(255) NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci