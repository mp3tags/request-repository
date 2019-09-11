CREATE TABLE IF NOT EXISTS request (
    `id` integer unsigned NOT NULL AUTO_INCREMENT,
    `created_at` TIMESTAMP not null default CURRENT_TIMESTAMP,
    `user_uuid` varchar(36) not null,
    `user_ip` varchar(255) not null,
    `url` varchar(255) not null,
    `data` varchar(255) not null,
     PRIMARY KEY (`id`)
) AUTO_INCREMENT=1;