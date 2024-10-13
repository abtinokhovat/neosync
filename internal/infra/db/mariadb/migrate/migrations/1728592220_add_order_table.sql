-- +migrate Up
create table `orders`(
    `id`  bigint unsigned not null primary key auto_increment,
    `created_at`   timestamp default current_timestamp not null,
    `updated_at`   timestamp default current_timestamp on update current_timestamp not null,
    `status` tinyint(1) not null default 1,
    `tracking_code` text not null,
    `customer_id`  bigint unsigned not null,
    `provider_id`  bigint unsigned not null,
    FOREIGN KEY (`customer_id`) REFERENCES `customers`(`id`),
    FOREIGN KEY (`provider_id`) REFERENCES `providers`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- +migrate Down
drop table `orders`;