-- name: CheckUserIsExisting :one
select exists(
    select 1 from users
    where username = sqlc.arg(username)
);

-- name: ListUsers :many
select
    username
from users
where removed_at is null;
