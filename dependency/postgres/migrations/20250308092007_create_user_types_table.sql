-- migrate:up
create table if not exists user_types (
    id serial,
    name varchar(255) not null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    primary key (id)
);

-- migrate:down
drop table if exists user_types;
