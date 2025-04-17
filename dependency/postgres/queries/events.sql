-- name: CreateEvent :copyfrom
insert into events (
    name, user_id
) values (
    $1, $2
);
