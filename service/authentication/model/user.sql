CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL DEFAULT '' COMMENT 'username like id',
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT 'Name of the user',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT 'User Password',
    `gender` ENUM ('male','female','other')  NOT NULL DEFAULT 'male' COMMENT 'Male | Female | Disclosed',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username_unique` (`username`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;