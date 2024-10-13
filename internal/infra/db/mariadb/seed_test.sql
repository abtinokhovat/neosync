-- providers
insert into `providers` (name, url) values
  ('mock-provider-1','https://localhost:9091/status'),
  ('mock-provider-2','https://localhost:9092/status');


-- customers
insert into `customers` (name, phone) values
    ('customer-a','09121111111'),
    ('customer-b','09121111112'),
    ('customer-c','09121111113');


insert into `orders` (`created_at`, `updated_at`, `status`, `tracking_code`, `customer_id`, `provider_id`)
values
    (NOW(), NOW(), 4, 'abc', 1, 1),
    (NOW(), NOW(), 2, 'jdk', 2, 1),
    (NOW(), NOW(), 3, 'TRACK12347', 3, 2),
    (NOW(), NOW(), 4, 'TRACK12348', 1, 2),
    (NOW(), NOW(), 3, 'lls', 2, 1),
    (NOW(), NOW(), 1, 'TRACK12350', 3, 2),
    (NOW(), NOW(), 2, 'sidf', 1, 1),
    (NOW(), NOW(), 3, 'asdf', 2, 1),
    (NOW(), NOW(), 3, 'TRACK12353', 3, 2),
    (NOW(), NOW(), 3, 'TRACK12354', 1, 2),
    (NOW(), NOW(), 4, 'weori', 2, 1),
    (NOW(), NOW(), 2, 'TRACK12356', 3, 2),
    (NOW(), NOW(), 3, 'awer', 1, 1),
    (NOW(), NOW(), 4, 'TRACK12358', 2, 2),
    (NOW(), NOW(), 3, 'dddc', 3, 1);