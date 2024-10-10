-- +migrate Up
create table `providers`(
    `id`  bigint unsigned not null primary key auto_increment,
    `name` text not null,
    `url` text not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

-- +migrate Down
drop table `providers`;