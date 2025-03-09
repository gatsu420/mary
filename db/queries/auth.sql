-- name: CheckUserIsExisting :one
select
    username is not null::bool as is_existing
from users
where username = sqlc.arg(username);
