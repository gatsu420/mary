-- migrate:up
create table if not exists food (
    id serial not null,
    name varchar(255) not null,
    type integer not null,
    intake_status integer not null,
    feeder integer not null,
    location integer not null,
    remarks varchar(255),
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    primary key (id)
);

-- migrate:down
drop table if exists food;
