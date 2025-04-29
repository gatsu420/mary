-- migrate:up
create table if not exists events (
    id serial,
    name varchar(255) not null,
    user_id varchar(255) not null,
    created_at timestamp with time zone,
    inserted_at timestamp with time zone default current_timestamp,
    primary key (id)
);

-- migrate:down
drop table if exists events;
