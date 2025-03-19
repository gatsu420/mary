-- name: CheckUserIsExisting :one
select exists (
    select 1 from users
    where username = sqlc.arg(username)
);
