-- +migrate Up
create table `order_status_history` (
  `id`  bigint unsigned not null primary key auto_increment,
  `order_id` bigint unsigned not null,
  `status` tinyint(1) not null,
  `changed_at` datetime not null,
  foreign key (order_id) references orders(id)
);

-- composite index on (order_id, status) for filtering by order and status
create index idx_order_id_status on order_status_history(order_id, status);


-- +migrate Down
drop table `order_status_history`;
drop index idx_order_id_status on order_status_history;