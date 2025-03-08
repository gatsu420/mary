-- migrate:up
create table if not exists users (
    id serial,
    name varchar(255) not null,
    type_id integer not null,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp,
    removed_at timestamp with time zone,
    primary key (id)
);

-- migrate:down
drop table if exists users;
