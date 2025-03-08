-- migrate:up
insert into user_types (name) values
    ('admin'),
    ('user');

-- migrate:down
truncate table user_types restart identity;
