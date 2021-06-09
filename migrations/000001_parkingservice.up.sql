CREATE TABLE `parking`
(
    `id`          VARCHAR(255),
    `name`       VARCHAR(255) NOT NULL,
    `description` TEXT         NOT NULL,
    `coordinates` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;