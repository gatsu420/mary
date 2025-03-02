-- migrate:up
alter table food
add column removed_at timestamp with time zone default null;

-- migrate:down
alter table food
drop column removed_at;
