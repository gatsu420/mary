-- name: CreateEvent :copyfrom
insert into events (
    name, user_id, created_at
) values (
    $1, $2, $3
);
